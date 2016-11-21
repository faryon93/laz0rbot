package botmux

// laz0rbot - telegram bot for llt
// Copyright (C) 2016 Maximilian Pachl

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

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