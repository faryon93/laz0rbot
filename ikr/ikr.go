package ikr

import (
    "log"
    "strings"

    "github.com/faryon93/laz0rbot/botmux"
)


// ----------------------------------------------------------------------------------
//  constants
// ----------------------------------------------------------------------------------

const (
    ATTEND_FINISH_TEXT = "You are now registrated as %s to IKR. I will send you a " +
                         "friendly reminder when IKR is about to start. You can " +
                         "leave IKR with */ikr leave* whenever you want."
)


// ----------------------------------------------------------------------------------
//  command muxer
// ----------------------------------------------------------------------------------

func Entry(ctx botmux.Context, args string) (botmux.CommandFunc) {
    switch strings.ToLower(args) {
        case "schedule":
            return Schedule(ctx, args)

        case "attend":
            return Attend(ctx, args)

        default:
            return Usage(ctx, args)
    }
}


// ----------------------------------------------------------------------------------
//  telegram commands
// ----------------------------------------------------------------------------------

func Usage(ctx botmux.Context, args string) (botmux.CommandFunc) {
    err := ctx.SendText("*IKR commands:*\n/ikr schedule - display schedule for the next six weeks\n/ikr attend - atted to IKR")
    if err != nil {
        log.Println("failed to send usage command")
    }

    return nil
}

func Schedule(ctx botmux.Context, args string) (botmux.CommandFunc) {
    // answer with a table containing the schedule
    err := ctx.SendText("*Schedule for the next six weeks:*\n // TODO")
    if err != nil {
        log.Println("failed to reply to schedule command:", err.Error())
    }

    return nil
}

func Attend(ctx botmux.Context, args string) (botmux.CommandFunc) {
    // ask the user for her/his participe type
    err := ctx.Ask("Do you want to attend IKR as a member or just be a silent observer?", []string{"Member", "Silent Observer"})
    if err != nil {
        log.Println("failed to reply to attend command:", err.Error())
    }

    return AttendFinish
}

func AttendFinish(ctx botmux.Context, args string) (botmux.CommandFunc) {
    err := ctx.SendText("You are now registrated as " + args + " to IKR. " + "I will send you a friendly reminder when IKR is about to start. You can leave IKR with */ikr leave* whenever you want.")
    if err != nil {
        log.Println("failed to send message:", err.Error())
    }

    return nil
}