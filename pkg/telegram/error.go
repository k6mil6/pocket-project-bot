package telegram

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	errInvalidURL   = errors.New("url is invalid")
	errUnauthorised = errors.New("user is not authorised")
	errUnableToSave = errors.New("unable to save")
)

func (b *Bot) handleError(chatID int64, err error) {
	msg := tgbotapi.NewMessage(chatID, b.messages.Default)

	switch {
	case errors.Is(err, errInvalidURL):
		msg.Text = b.messages.InvalidURL
		b.bot.Send(msg)
	case errors.Is(err, errUnauthorised):
		msg.Text = b.messages.Unauthorised
		b.bot.Send(msg)
	case errors.Is(err, errUnableToSave):
		msg.Text = b.messages.UnableToSave
		b.bot.Send(msg)
	default:
		b.bot.Send(msg)
	}
}
