package main

import (
	"fmt"
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

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_KEY"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	yandexClient := yandex.NewClient(os.Getenv("YANDEX_URL"), os.Getenv("YANDEX_KEY"))
	useCase := translation.NewUseCase(yandexClient)
	defaultCommand := _default.NewCommand(useCase)
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if m := update.Message; m != nil { // If we got a message
			var msg tgbotapi.MessageConfig

			if m.IsCommand() {
				// commands
				continue
			}

			msg, err = defaultCommand.Handle(*m)
			if err != nil {
				fmt.Println(err)
			}

			bot.Send(msg)
		}
	}
}
