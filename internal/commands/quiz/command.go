package quiz

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"githab.com/oleglacto/translater/internal/pkg/quiz"
)

type Quiz interface {
	Start(chatID int) error
	End(chatID int)
	GetQuestion(chatID int) (quiz.IrregularVerbsGameStruct, error)
	//CheckAnswer(chatID int) bool
}

type Handler struct {
	quiz Quiz
	bot  *tgbotapi.BotAPI
}

func New(quiz Quiz, bot *tgbotapi.BotAPI) *Handler {
	return &Handler{quiz: quiz, bot: bot}
}

func (h Handler) Handle(ctx context.Context, message tgbotapi.Message) ([]tgbotapi.MessageConfig, error) {
	if message.IsCommand() {
		err := h.quiz.Start(int(message.Chat.ID))
		if err != nil {
			return nil, err
		}

		question, err := h.quiz.GetQuestion(int(message.Chat.ID))
		if err != nil {
			return nil, err
		}

		messages := []tgbotapi.MessageConfig{{Text: "Выбери правильный ответ: "}}

		poll := tgbotapi.NewPoll(message.Chat.ID, question.Question, question.AnswerOptions...)
		poll.Type = "quiz"
		poll.CorrectOptionID = int64(question.CorrectAnswerIndex)
		h.bot.Send(poll)

		return messages, nil
	}
	return nil, nil
}
