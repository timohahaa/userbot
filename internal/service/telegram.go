package service

import (
	"context"
	"errors"
	"math/rand/v2"

	"github.com/timohahaa/userbot/internal/entity"
	"github.com/timohahaa/userbot/internal/repository"

	"github.com/celestix/gotgproto"
	"github.com/celestix/gotgproto/sessionMaker"
	"github.com/gotd/td/tg"
	log "github.com/sirupsen/logrus"
)

type Client = gotgproto.Client

var PREDEFINED_REACTIONS = []tg.ReactionClass{
	&tg.ReactionEmoji{Emoticon: "👍"},
	&tg.ReactionEmoji{Emoticon: "🎉"},
	&tg.ReactionEmoji{Emoticon: "💯"},
	&tg.ReactionEmoji{Emoticon: "👎"},
	&tg.ReactionEmoji{Emoticon: "🔥"},
	&tg.ReactionEmoji{Emoticon: "🤩"},
	&tg.ReactionEmoji{Emoticon: "❤️‍🔥"},
	&tg.ReactionEmoji{Emoticon: "❤️"},
	&tg.ReactionEmoji{Emoticon: "😁"},
}

type telegramService struct {
	apiId            int
	apiHash          string
	repo             repository.TelegramRepository
	commentFrequency int
}

func NewTelegramService(
	repo repository.TelegramRepository,
	apiId int,
	apiHash string,
	commentFrequency int,
) *telegramService {
	return &telegramService{
		apiId:            apiId,
		apiHash:          apiHash,
		repo:             repo,
		commentFrequency: commentFrequency,
	}
}

func (s *telegramService) GetRandomAccount(ctx context.Context) (entity.Account, error) {
	account, err := s.repo.GetRandomAccount(ctx)

	if err != nil {
		return entity.Account{}, err
	}

	return account, nil
}

func (s *telegramService) SaveChannelByName(ctx context.Context, channelName string) error {
	// resolve the channel by name (or link like https://t.me/channelname)
	// save the channel struct to db, so it can be later accessed by id and access hash (thats more stable)

	client, err := s.getRandomUserClient(ctx)
	if err != nil {
		return err
	}
	defer client.Stop()

	resolvedPeer, err := client.API().ContactsResolveUsername(ctx, channelName)
	if err != nil {
		return err
	}

	channel, ok := resolvedPeer.Chats[0].(*tg.Channel)
	if !ok {
		return ErrChannelNotFound
	}

	// save the channel
	chanEntity := entity.Channel{
		Id:         channel.ID,
		AccessHash: channel.AccessHash,
		UserCount:  channel.ParticipantsCount,
		Name:       channel.Username,
	}

	err = s.repo.SaveChannel(ctx, chanEntity)
	if err != nil {
		return ErrFailedToSaveChannel
	}

	return nil
}

func (s *telegramService) getChannelByChannelId(ctx context.Context, id int64) (entity.Channel, error) {
	channel, err := s.repo.GetChannelByChannelId(ctx, id)
	if err != nil {
		return entity.Channel{}, err
	}
	return channel, nil
}

func (s *telegramService) getRandomUserClient(ctx context.Context) (*gotgproto.Client, error) {
	randomAccount, err := s.GetRandomAccount(ctx)
	if err != nil {
		return nil, err
	}

	clientType := gotgproto.ClientType{
		Phone: randomAccount.PhoneNumber,
	}

	client, err := gotgproto.NewClient(
		s.apiId,
		s.apiHash,
		clientType,
		&gotgproto.ClientOpts{
			DisableCopyright: true,
			Session:          sessionMaker.StringSession(randomAccount.SessionString),
		},
	)
	if err != nil {
		return nil, err
	}

	return client, err
}

func (s *telegramService) getAvailableReactions(ctx context.Context, client *Client, channel entity.Channel) ([]tg.ReactionClass, error) {
	chat, err := client.API().ChannelsGetFullChannel(ctx, &tg.InputChannel{
		ChannelID:  channel.Id,
		AccessHash: channel.AccessHash,
	})
	if err != nil {
		return nil, err
	}

	fullChannel := chat.FullChat.(*tg.ChannelFull)
	// reactions can be set by admin or all reactions can be allowed
	// in case that all of them are allowed - return a set of predefined non-premium reactions
	switch t := fullChannel.AvailableReactions.(type) {
	case *tg.ChatReactionsSome:
		// admin-set reactions
		return t.Reactions, nil
	case *tg.ChatReactionsAll:
		// all reactions are allowed
		return PREDEFINED_REACTIONS, nil
	default:
		return nil, errors.New("should never return an error from here???")
	}
}

func (s *telegramService) getLastChannelPostId(ctx context.Context, client *Client, channel entity.Channel) (int, error) {
	msgs, err := client.API().MessagesGetHistory(ctx, &tg.MessagesGetHistoryRequest{
		Peer: &tg.InputPeerChannel{
			ChannelID:  channel.Id,
			AccessHash: channel.AccessHash,
		},
		Limit: 1,
	})
	if err != nil {
		return 0, err
	}

	chanMessages := msgs.(*tg.MessagesChannelMessages)
	lastPost := chanMessages.Messages[0]
	return lastPost.GetID(), nil
}

func (s *telegramService) ReactNewPost(ctx context.Context, channelId int64) error {
	// get the client to perform all preparation - reactions, last post, etc
	client, err := s.getRandomUserClient(ctx)
	if err != nil {
		return err
	}

	// get the channel first
	channel, err := s.getChannelByChannelId(ctx, channelId)
	if err != nil {
		return ErrChannelNotFound
	}

	// calculate how many times to comment/react to a new post
	reactionCount := channel.UserCount / s.commentFrequency

	// get available reactions (list of admin-set reactions or a list of default ones)
	reactions, err := s.getAvailableReactions(ctx, client, channel)
	if err != nil {
		return err
	}

	// get last post id (the one to react to)
	postId, err := s.getLastChannelPostId(ctx, client, channel)
	if err != nil {
		return err
	}

	// stop the preparation client
	client.Stop()

	for _ = range reactionCount {
		reactionIdx := rand.IntN(len(reactions))
		randomUser, err := s.getRandomUserClient(ctx)
		if err != nil {
			log.Errorf("error getting random user client", err)
			continue
		}

		_, err = randomUser.API().MessagesSendReaction(context.Background(), &tg.MessagesSendReactionRequest{
			Peer: &tg.InputPeerChannel{
				ChannelID:  channel.Id,
				AccessHash: channel.AccessHash,
			},
			MsgID:    postId,
			Reaction: []tg.ReactionClass{reactions[reactionIdx]},
		})
		if err != nil {
			log.Errorf("error reacting to post", err)
			continue
		}
	}

	return nil
}
