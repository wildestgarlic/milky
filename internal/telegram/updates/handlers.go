package updates

import (
	"TelebotOne/internal/config"
	"TelebotOne/internal/repo/exersises"
	"TelebotOne/internal/telegram/constant"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

var cfg = config.Config

func (b *Bot) SendMessage(message *tgbotapi.Message, text string) {
	msg := tgbotapi.NewMessage(message.Chat.ID, text)

	_, err := b.bot.Send(msg)
	if err != nil {
		log.Printf(">sendMessage error: %v", err) //TODO: handle error
	}
}

func (b *Bot) handleCommands(message *tgbotapi.Message) {
	switch message.Command() {

	case constant.StartCommand:
		b.handleStart(message)
		return
	case constant.CreateCommand:
		b.handleCreate(message)
		return
	case constant.DropCommand:
		b.handleDrop(message)
		return
	case constant.ShowAllCommand:
		b.handleShowAll(message)
		return
	}
}

func (b *Bot) handleMessages(message *tgbotapi.Message) {
	if cfg.BotDebug { //fixme
		log.Printf("[%s] %s", message.From.UserName, message.Text)
	}

	b.SendMessage(message, message.Text) //fixme: echo logic
}

func (b *Bot) handleStart(message *tgbotapi.Message) {
	text := constant.StartMessage
	b.SendMessage(message, text)
}

func (b *Bot) handleCreate(message *tgbotapi.Message) {
	repo := exersises.NewExercisesRepo()

	text := constant.CreateSuccessMessage

	err := repo.CreateExerciseTable("exercises") //fixme: fix name
	if err != nil {
		text = constant.CreateErrorMessage
		if err.Error() == constant.TableExist {
			text = constant.AlreadyExistErrorMessage
		}
	}

	b.SendMessage(message, text)
}

func (b *Bot) handleDrop(message *tgbotapi.Message) {
	text := constant.DropSuccessMessage
	b.SendMessage(message, text)
}

func (b *Bot) handleShowAll(message *tgbotapi.Message) {
	text := constant.ShowAllSuccessMessage
	b.SendMessage(message, text)
}
