package telegram

import "errors"

var (
	ErrChannelNotFound     = errors.New("channel not found")
	ErrFailedToSaveChannel = errors.New("failed to save channel")
)
