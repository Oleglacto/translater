package _default

import (
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

func NewCommand(translation Translation) *Command {
	return &Command{translation: translation}
}

func (c Command) Handle(message tgbotapi.Message) (tgbotapi.MessageConfig, error) {
	result, err := c.translation.Translate(message.Text)
	if err != nil {
		return tgbotapi.MessageConfig{}, err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, result)

	return msg, nil
}
