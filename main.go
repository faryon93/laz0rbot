package main

import (
    "log"
    "os"

    "github.com/faryon93/laz0rbot/botmux"
    "github.com/faryon93/laz0rbot/ikr"

    "gopkg.in/telegram-bot-api.v4"
)


// ----------------------------------------------------------------------------------
//  application entry
// ----------------------------------------------------------------------------------

func main() {
    // register with the telegram bot api
    err := botmux.Register(os.Args[1])
    if err != nil {
        log.Println("failed to register with bot")
        os.Exit(-1)
    }
    log.Printf("registered with telegram bot @%s", botmux.Bot.Self.UserName)

    // register bot commands
    botmux.Command("/start", BotStart)
    botmux.Command("/ikr", ikr.Entry)

    // process all incoming commands
    err = botmux.Listen()
    if err != nil {
        log.Println("failed to listen for updates:", err.Error())
    }
}


// ----------------------------------------------------------------------------------
//  telegram commands
// ----------------------------------------------------------------------------------

func BotStart(ctx botmux.Context, args string) (botmux.CommandFunc) {
    msg := tgbotapi.NewMessage(ctx.Message.Chat.ID, "Welcome to the LLT telegram bot! How can I help you?\n\nAvailable Commands:\n/ikr - IKR planning")
    ctx.Bot.Send(msg)

    return nil
}
