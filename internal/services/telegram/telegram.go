package telegram

import (
	"context"

	"github.com/timohahaa/userbot/internal/entity"
	"github.com/timohahaa/userbot/internal/repository"

	"github.com/celestix/gotgproto"
	"github.com/celestix/gotgproto/sessionMaker"
	"github.com/gotd/td/tg"
)

type telegramService struct {
	apiId       int
	apiHash     string
	accountRepo repository.AccountRepository
	channelRepo repository.ChannelRepository
}

func NewTelegramService(
	accRepo repository.AccountRepository,
	chanRepo repository.ChannelRepository,
	apiId int,
	apiHash string,
) *telegramService {
	return &telegramService{}
}

func (s *telegramService) GetRandomAccount(ctx context.Context) (entity.Account, error) {
	account, err := s.accountRepo.GetRandomAccount(ctx)

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

	resolvedPeer, err := client.API().ContactsResolveUsername(context.Background(), channelName)
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

	err = s.channelRepo.SaveChannel(ctx, chanEntity)
	if err != nil {
		return ErrFailedToSaveChannel
	}

	return nil
}

func (s *telegramService) getChannelByChannelId(ctx context.Context, id int64) (entity.Channel, error) {
	channel, err := s.channelRepo.GetChannelByChannelId(ctx, id)
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

func (s *telegramService) CommentNewPost(ctx context.Context, channelId int64) error {
	// get the channel first
	channel, err := s.getChannelByChannelId(ctx, channelId)
	if err != nil {
		return ErrChannelNotFound
	}

	client, err := s.getRandomUserClient(ctx)
	if err != nil {
		return err
	}

	client.API().MessagesGetHistory(ctx, &tg.MessagesGetHistoryRequest{})

	return nil
}
