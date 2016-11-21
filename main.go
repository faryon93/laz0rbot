package main

import (
    "log"
    "os"

    "github.com/faryon93/laz0rbot/botmux"
    "github.com/faryon93/laz0rbot/ikr"
    "github.com/faryon93/laz0rbot/state"
)

// ----------------------------------------------------------------------------------
//  constants
// ----------------------------------------------------------------------------------

const (
    TEXT_USAGE = "Welcome to the LLT telegram bot! How can I help you?\n\n" +
                 "Available Commands:\n" +
                 "/ikr - IKR planning\n" + 
                 "/name - set your display name"
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
    botmux.Command("/start", Usage)
    botmux.Command("/usage", Usage)
    botmux.Command("/help", Usage)
    botmux.Command("/ikr", ikr.Entry)
    botmux.Command("/name", Name)

    // process all incoming commands
    err = botmux.Listen()
    if err != nil {
        log.Println("failed to listen for updates:", err.Error())
    }
}


// ----------------------------------------------------------------------------------
//  telegram commands
// ----------------------------------------------------------------------------------

func Usage(ctx botmux.Context, args string) (botmux.CommandFunc) {
    err := ctx.SendText(TEXT_USAGE)
    if err != nil {
        log.Println("failed to send text:", err.Error())
    }

    return nil
}

func Name(ctx botmux.Context, args string) (botmux.CommandFunc) {
    err := ctx.SendText("How should we call you?")
    if err != nil {
        log.Println("failed to send text:", err.Error())
    }

    return NameFinish
}

func NameFinish(ctx botmux.Context, args string) (botmux.CommandFunc) {
    err := ctx.SendText("Okay, from now on I will call you %s.", args)
    if err != nil {
        log.Println("failed to send text:", err.Error())
    }

    state.State.Aliases[ctx.Message.From.ID] = args
    state.State.Save()

    return nil
}
