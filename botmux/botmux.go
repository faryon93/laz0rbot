package botmux

import (
    "errors"
    "log"
    "strings"

    "gopkg.in/telegram-bot-api.v4"
)


// ----------------------------------------------------------------------------------
//  constants
// ----------------------------------------------------------------------------------

const (
    TELEGRAM_UPDATE_TIMEOUT = 60
)


// ----------------------------------------------------------------------------------
//  types
// ----------------------------------------------------------------------------------

type CommandFunc func(ctx Context, args string) (CommandFunc)


// ----------------------------------------------------------------------------------
//  global variables
// ----------------------------------------------------------------------------------

var Bot *tgbotapi.BotAPI

var commands map[string]CommandFunc = make(map[string]CommandFunc)
var sessions map[int]CommandFunc = make(map[int]CommandFunc)


// ----------------------------------------------------------------------------------
//  functions
// ----------------------------------------------------------------------------------

func Register(token string) (error) {
    var err error

    // register with the telegram bot api
    Bot, err = tgbotapi.NewBotAPI(token)
    return err
}

func Command(name string, handler CommandFunc) {
    commands[name] = handler
}

func Listen() (error) {
    if Bot == nil {
        return errors.New("not registered with telegram bot api")
    }

    // get an update channel
    u := tgbotapi.NewUpdate(0)
    u.Timeout = TELEGRAM_UPDATE_TIMEOUT
    updates, err := Bot.GetUpdatesChan(u)
    if err != nil {
        return err
    }

    // process all updates received by the bot
    for update := range updates {
        // we received a private message
        if update.Message != nil {
            log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

            // a new command was received
            if (strings.HasPrefix(update.Message.Text, "/")) {
                args := strings.SplitN(update.Message.Text, " ", 2)
                if (len(args) < 1) {
                    continue
                }

                // find a suitable command handler
                processed := false
                for command, handler := range commands {
                    if (command == args[0]) {
                        // execute the handler function and update the user session
                        ctx := Context{
                            Bot: Bot,
                            Message: update.Message,
                        }
                        sessions[update.Message.From.ID] = handler(ctx, args[1])
                        
                        // finished command handler
                        processed = true
                        break
                    }
                }

                // no apropriate command handler was found
                if !processed {
                    log.Println("could not find command handler for", args[0])
                }

            // a plain message was received -> preceed with the user session
            } else {
                // check if a valid session is already registrated
                handler, valid := sessions[update.Message.From.ID]
                if valid && handler != nil {
                    ctx := Context{
                            Bot: Bot,
                            Message: update.Message,
                    }
                    sessions[update.Message.From.ID] = handler(ctx, update.Message.Text)

                // no valid session was found
                } else {
                    log.Printf("[%s] received plain message, without user session\n", update.Message.From.UserName)
                }
            }
        }
    }

    return nil
}