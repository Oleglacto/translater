package irregular_verb

import (
	"context"
	"fmt"
	"githab.com/oleglacto/translater/internal/pkg/helpers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"githab.com/oleglacto/translater/internal/pkg/language"
)

type RandomIrregularVerbGetter interface {
	GetRandomVerb() (language.IrregularVerb, error)
}

type Handler struct {
	irregularVerbGetter RandomIrregularVerbGetter
}

func New(getter RandomIrregularVerbGetter) *Handler {
	return &Handler{irregularVerbGetter: getter}
}

func (h Handler) Handle(ctx context.Context, message tgbotapi.Message) ([]tgbotapi.MessageConfig, error) {
	verb, err := h.irregularVerbGetter.GetRandomVerb()
	if err != nil {
		return nil, err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(
		"%s — %s — %s — ||%s||",
		helpers.EscapeBrackets(verb.Infinitive),
		helpers.EscapeBrackets(verb.V2),
		helpers.EscapeBrackets(verb.V3),
		helpers.EscapeBrackets(verb.Translate),
	))
	msg.ParseMode = tgbotapi.ModeMarkdownV2

	return []tgbotapi.MessageConfig{msg}, nil
}
