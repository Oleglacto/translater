package server

import (
	"context"
	"errors"
	"fmt"
	"githab.com/oleglacto/translater/internal/pkg/cache"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var unknownCommandErr = errors.New("unknown command")

type handleFunc func(ctx context.Context, msg tgbotapi.Message) ([]tgbotapi.MessageConfig, error)

type Server struct {
	bot *tgbotapi.BotAPI

	cache  *cache.Cache
	routes map[string]handleFunc
}

func New(c *cache.Cache, b *tgbotapi.BotAPI) *Server {
	routes := make(map[string]handleFunc)
	return &Server{
		bot:    b,
		routes: routes,
		cache:  c,
	}
}

func (s *Server) RegisterCommand(name string, handler func(ctx context.Context, msg tgbotapi.Message) ([]tgbotapi.MessageConfig, error)) {
	s.routes[name] = handler
}

func (s *Server) Run() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := s.bot.GetUpdatesChan(u)

	for update := range updates {
		if m := update.Message; m != nil {
			messages, err := s.handle(*m)
			if err != nil {
				log.Println(fmt.Errorf("error: %w", err).Error())
				switch {
				case errors.Is(err, unknownCommandErr):
					s.bot.Send(tgbotapi.MessageConfig{Text: "Неизвестная команда :hmm:", BaseChat: tgbotapi.BaseChat{ChatID: m.Chat.ID}})
				default:
					s.bot.Send(tgbotapi.MessageConfig{Text: "Извини, произошла ошибка :( Попробуй позже", BaseChat: tgbotapi.BaseChat{ChatID: m.Chat.ID}})
				}
				continue
			}
			for _, msg := range messages {
				msg.ChatID = m.Chat.ID
				s.bot.Send(msg)
			}
		}
	}
}

func (s Server) handle(message tgbotapi.Message) ([]tgbotapi.MessageConfig, error) {
	if message.IsCommand() {
		command, ok := s.routes[message.Command()]
		if !ok {
			return nil, unknownCommandErr
		}
		return command(context.Background(), message)
	}

	if message.Text == "" {
		return nil, nil
	}

	return s.routes["default"](context.Background(), message)
}
