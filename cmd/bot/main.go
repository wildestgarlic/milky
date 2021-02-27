package main

import (
	"TelebotOne/internal/telegram"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

func main() {
	bot := getBot()
	bot.Debug = true

	startBot(bot)
}

func getBot() *tgbotapi.BotAPI {
	token := os.Getenv("BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Start connection error: %s", err)
	}

	return bot
}

func startBot(bot *tgbotapi.BotAPI) {
	telegramBot := telegram.NewBot(bot)
	err := telegramBot.Start()
	if err != nil {
		log.Fatalf("Bot initializing error: %s", err)
	}
}
