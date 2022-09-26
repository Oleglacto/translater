package random

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h Handler) Handle(ctx context.Context, message tgbotapi.Message) ([]tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "random")
	return []tgbotapi.MessageConfig{msg, msg}, nil
}
