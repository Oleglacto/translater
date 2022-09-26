package _default

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Translation interface {
	Translate(text string) (string, error)
}

type State interface {
}

type Command struct {
	translation Translation
	userState   State
}

func New(translation Translation) *Command {
	return &Command{translation: translation}
}

func (c Command) Handle(ctx context.Context, message tgbotapi.Message) ([]tgbotapi.MessageConfig, error) {
	var messages []tgbotapi.MessageConfig
	result, err := c.translation.Translate(message.Text)
	if err != nil {
		return nil, err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, result)
	messages = append(messages, msg)
	return messages, nil
}
