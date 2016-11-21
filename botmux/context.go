package botmux

import (
	"fmt"

	"gopkg.in/telegram-bot-api.v4"
)


// ----------------------------------------------------------------------------------
//  types
// ----------------------------------------------------------------------------------

type Context struct {
	Bot *tgbotapi.BotAPI
	Message *tgbotapi.Message
}


// ----------------------------------------------------------------------------------
//  member functions
// ----------------------------------------------------------------------------------

func (this *Context) SendText(format string, values ...interface{}) (error) {
	// setup the text
	message := fmt.Sprintf(format, values...)

	// send a text message to the user
	msg := tgbotapi.NewMessage(this.Message.Chat.ID, message)
	msg.ParseMode = "Markdown"
    _, err := this.Bot.Send(msg)

    return err
}

func (this *Context) Ask(question string, buttons []string) (error) {
	// construct the question message
	msg := tgbotapi.NewMessage(this.Message.Chat.ID, question)
	msg.ParseMode = "Markdown"

	// setup predefined answers
	keys := make([]tgbotapi.KeyboardButton, len(buttons))
	for i, text := range buttons {
		keys[i] = tgbotapi.NewKeyboardButton(text)
	}

	// setup the keyboard
	keyboard := tgbotapi.NewReplyKeyboard(keys);
	keyboard.OneTimeKeyboard = true
	keyboard.Selective = true
	msg.ReplyMarkup = keyboard

	// send the question to the user
    _, err := this.Bot.Send(msg)

    return err
}