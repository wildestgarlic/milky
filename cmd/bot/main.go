package main

import (
	"TelebotOne/internal/config"
	"TelebotOne/internal/telegram/updates"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

var cfg = config.Config

func main() {
	bot := getBot()

	debugToggle(bot)

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
	telegramBot := updates.NewBot(bot)
	err := telegramBot.Start()
	if err != nil {
		log.Fatalf("Bot initializing error: %s", err)
	}
}

func debugToggle(bot *tgbotapi.BotAPI) {
	if cfg.BotDebug {
		bot.Debug = true
	}
}
