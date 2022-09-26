package main

import (
	"fmt"
	"githab.com/oleglacto/translater/internal/commands/irregular_verb"
	quizCommand "githab.com/oleglacto/translater/internal/commands/quiz"
	"githab.com/oleglacto/translater/internal/commands/random"
	"githab.com/oleglacto/translater/internal/pkg/cache"
	"githab.com/oleglacto/translater/internal/pkg/language"
	"githab.com/oleglacto/translater/internal/pkg/quiz"
	"githab.com/oleglacto/translater/internal/pkg/router"
	"githab.com/oleglacto/translater/internal/usecase/random_irregular_verb"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"

	_default "githab.com/oleglacto/translater/internal/commands/default"
	"githab.com/oleglacto/translater/internal/pkg/yandex"
	"githab.com/oleglacto/translater/internal/usecase/translation"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}

	// bot init
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_KEY"))
	if err != nil {
		log.Panic(err)
	}

	// app
	c := cache.New()
	r := router.New(c)
	yandexClient := yandex.NewClient(os.Getenv("YANDEX_URL"), os.Getenv("YANDEX_KEY"))
	irregularVerbsGetter := language.NewIrregularVerbsRepository()
	useCaseTranslate := translation.NewUseCase(yandexClient)
	useCaseRandomIrregularVerb := random_irregular_verb.New(irregularVerbsGetter)
	irregularVerbQuiz := quiz.NewIrregularVerbs(irregularVerbsGetter, c)

	r.Add("default", false, _default.New(useCaseTranslate).Handle)
	r.Add("random", true, random.New().Handle)
	r.Add("irregular_verb", true, irregular_verb.New(useCaseRandomIrregularVerb).Handle)
	r.Add("quiz", true, quizCommand.New(irregularVerbQuiz, bot).Handle)

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if m := update.Message; m != nil {
			messages, err := r.Do(*m)
			if err != nil {
				log.Println(err)
				bot.Send(tgbotapi.MessageConfig{Text: "Извини, произошла ошибка :( Попробуй позже", BaseChat: tgbotapi.BaseChat{ChatID: m.Chat.ID}})
				continue
			}
			for _, msg := range messages {
				msg.ChatID = m.Chat.ID
				bot.Send(msg)
			}
		}

		if callback := update.CallbackQuery; callback != nil {
			fmt.Println(callback)
			callback2 := tgbotapi.NewCallbackWithAlert(update.CallbackQuery.ID, "\\xF0\\x9F\\x98\\x81")
			if _, err := bot.Request(callback2); err != nil {
				panic(err)
			}

			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}

		if answer := update.PollAnswer; answer != nil {
			fmt.Println(answer)
		}
	}
}
