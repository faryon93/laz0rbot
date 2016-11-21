package botmux

import (
	"gopkg.in/telegram-bot-api.v4"
)


// ----------------------------------------------------------------------------------
//  types
// ----------------------------------------------------------------------------------

type Context struct {
	Bot *tgbotapi.BotAPI
	Message *tgbotapi.Message
}