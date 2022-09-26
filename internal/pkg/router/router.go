package router

import (
	"context"
	"errors"
	"fmt"
	"githab.com/oleglacto/translater/internal/pkg/cache"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type handleFunc func(ctx context.Context, msg tgbotapi.Message) ([]tgbotapi.MessageConfig, error)

type MessageRouter struct {
	cache  *cache.Cache
	routes map[string]handleFunc
}

func New(cache *cache.Cache) *MessageRouter {
	routes := make(map[string]handleFunc)
	return &MessageRouter{cache: cache, routes: routes}
}

func (r MessageRouter) Do(msg tgbotapi.Message) ([]tgbotapi.MessageConfig, error) {
	key := "default"

	if msg.IsCommand() {
		key = commandName(msg.Command())
	}

	quizStarted := false
	_, err := r.cache.Get(fmt.Sprintf("quiz.irregular_verbs.chat_id.%d", msg.Chat.ID))
	if err != nil {
		if !errors.Is(err, cache.NotFound) {
			return nil, err
		}
	}
	if err == nil {
		quizStarted = true
	}

	if quizStarted {
		key = "quiz"
	}

	handler, ok := r.routes[key]
	if !ok {
		return []tgbotapi.MessageConfig{}, errors.New("no handler")
	}

	return handler(context.Background(), msg)
}

func (r *MessageRouter) Add(name string, isCommand bool, handler func(ctx context.Context, msg tgbotapi.Message) ([]tgbotapi.MessageConfig, error)) {
	if isCommand {
		r.routes[commandName(name)] = handler
		return
	}

	r.routes[name] = handler
}

func commandName(name string) string {
	return "command_" + name
}
