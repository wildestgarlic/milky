package main

import (
	"TelebotOne/internal/telegram"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

func main() {
	var err error

	token := os.Getenv("TOKEN")
	fmt.Println(token)
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot)
	err = telegramBot.Start()
	if err != nil {
		log.Fatal(err)
	}

}

