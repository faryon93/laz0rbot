package main

import (
    "log"
    "os"

    "github.com/faryon93/laz0rbot/botmux"
    "github.com/faryon93/laz0rbot/ikr"
)

// ----------------------------------------------------------------------------------
//  constants
// ----------------------------------------------------------------------------------

const (
    TEXT_USAGE = "Welcome to the LLT telegram bot! How can I help you?\n\n" +
                 "Available Commands:\n" +
                 "/ikr - IKR planning"
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
    err := ctx.SendText(TEXT_USAGE)
    if err != nil {
        log.Println("failed to send text:", err.Error())
    }

    return nil
}
