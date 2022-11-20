package main

import (
	"githab.com/oleglacto/translater/internal/commands/irregular_verb"
	quizCommand "githab.com/oleglacto/translater/internal/commands/quiz"
	"githab.com/oleglacto/translater/internal/commands/random"
	"githab.com/oleglacto/translater/internal/pkg/cache"
	"githab.com/oleglacto/translater/internal/pkg/language"
	"githab.com/oleglacto/translater/internal/pkg/quiz"
	"githab.com/oleglacto/translater/internal/pkg/server"
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
	yandexClient := yandex.NewClient(os.Getenv("YANDEX_URL"), os.Getenv("YANDEX_KEY"))
	irregularVerbsGetter := language.NewIrregularVerbsRepository()
	useCaseTranslate := translation.NewUseCase(yandexClient)
	useCaseRandomIrregularVerb := random_irregular_verb.New(irregularVerbsGetter)
	irregularVerbQuiz := quiz.NewIrregularVerbs(irregularVerbsGetter, c)

	bot.Debug = true
	s := server.New(c, bot)
	s.RegisterCommand("default", _default.New(useCaseTranslate).Handle)
	s.RegisterCommand("random", random.New().Handle)
	s.RegisterCommand("irregular_verb", irregular_verb.New(useCaseRandomIrregularVerb).Handle)
	s.RegisterCommand("quiz", quizCommand.New(irregularVerbQuiz, bot).Handle)

	s.Run()
}
