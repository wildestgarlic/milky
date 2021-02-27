package telegram

import (
	"TelebotOne/internal/telegram/constant"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) {
	switch message.Command() {
	case constant.Start:
		b.handleStart(message)
		return
	case constant.CreateExercise:
		b.handleCreate(message)
		return
	case constant.DropExercise:
		b.handleDrop(message)
		return
	case constant.ShowAllExercises:
		b.handleShowAll(message)
		return
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	_, err := b.bot.Send(msg)
	if err != nil {
		log.Printf("handleMessage >sendMessage error: %v", err) //TODO: handle error
	}
}

func (b *Bot) handleStart(message *tgbotapi.Message) {

}

func (b *Bot) handleCreate(message *tgbotapi.Message) {

}

func (b *Bot) handleDrop(message *tgbotapi.Message) {

}

func (b *Bot) handleShowAll(message *tgbotapi.Message) {

}
