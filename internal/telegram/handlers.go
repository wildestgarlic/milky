package telegram

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) {
	switch message.Command() {

	}
}

func (b *Bot) handleMassage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	b.bot.Send(msg)
}
