package ikr

import (
    "log"
    "strings"

    "github.com/faryon93/laz0rbot/botmux"
    "github.com/faryon93/laz0rbot/state"
)


// ----------------------------------------------------------------------------------
//  constants
// ----------------------------------------------------------------------------------

const (
    TEXT_USAGE =          "*IKR commands:*\n" +
                          "/ikr schedule - display schedule for the next six weeks\n" +
                          "/ikr attend - attend to IKR meetings\n" +
                          "/ikr leave - leave IKR\n" +
                          "/ikr delay - delay the next scheduled IKR meeting\n" + 
                          "/ikr pitchin - pitch in for the attende of todays meeting"
    TEXT_ATTEND =         "Do you want to attend IKR as a member or just be a silent observer?"
    TEXT_ATTEND_FINISH  = "You are now registrated as %s to IKR. I will send you a " +
                          "friendly reminder when IKR is about to start. You can " +
                          "leave IKR with */ikr leave* whenever you want."
    TEXT_ATTEND_ALREADY = "You are already attending IKR meetings."
    TEXT_DELAY =          "How long should the next IKR be delayed?"
    TEXT_DELAY_FINISHED = "The next IKR meeting has been re-scheduled on %s ."
    TEXT_PITCHIN =        "You have been registrated as attende of todays meeting. " +
                          "Tank you for your service!"
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

        case "delay":
            return Delay(ctx, args)

        case "pitchin":
            return PitchIn(ctx, args)

        default:
            return Usage(ctx, args)
    }
}


// ----------------------------------------------------------------------------------
//  telegram commands
// ----------------------------------------------------------------------------------

func Usage(ctx botmux.Context, args string) (botmux.CommandFunc) {
    err := ctx.Ask(TEXT_USAGE, []string{"/ikr schedule", "/ikr delay", "/ikr pitchin"})
    if err != nil {
        log.Println("failed to reply to attend command:", err.Error())
    }


    return nil
}

// ----------------------------------------------------------------------------------
func Schedule(ctx botmux.Context, args string) (botmux.CommandFunc) {
    // answer with a table containing the schedule
    err := ctx.SendText("*Schedule for the next six weeks:*\n // TODO")
    if err != nil {
        log.Println("failed to reply to schedule command:", err.Error())
    }

    return nil
}

// ----------------------------------------------------------------------------------
func Attend(ctx botmux.Context, args string) (botmux.CommandFunc) {
    // check if the user has already registrated for ikr module
    if state.State.Ikr.Attendees.Contains(ctx.Message.From.ID) {
        err := ctx.SendText(TEXT_ATTEND_ALREADY)
        if err != nil {
            log.Println("failed to reply to attend command:", err.Error())
        }

        return nil

    // ask the user for her/his attendance type
    } else {
        err := ctx.Ask(TEXT_ATTEND, []string{"Member", "Silent Observer"})
        if err != nil {
            log.Println("failed to reply to attend command:", err.Error())
        }
        
        return AttendFinish
    }
}

func AttendFinish(ctx botmux.Context, args string) (botmux.CommandFunc) {
    // check if a valid subscirption type is sent
    if args != "Member" && args != "Silent Observer" {
        err := ctx.SendText("Invalid subscription type: %s", args)
        if err != nil {
            log.Println("failed to send message:", err.Error())
        }
        return nil
    }

    // add the requesting user to the attendee list
    state.State.AddIkrAttendee(ctx.Message.From.ID, args)
    state.State.Save()

    // everything is fine -> send user a message
    err := ctx.SendText(TEXT_ATTEND_FINISH, args)
    if err != nil {
        log.Println("failed to send message:", err.Error())
    }

    return nil
}

// ----------------------------------------------------------------------------------
func Delay(ctx botmux.Context, args string) (botmux.CommandFunc) {
    err := ctx.SendText(TEXT_DELAY)
    if err != nil {
        log.Println("failed to send message:", err.Error())
    }

    return DelayFinished   
}

func DelayFinished(ctx botmux.Context, args string) (botmux.CommandFunc) {
    err := ctx.SendText(TEXT_DELAY_FINISHED, "2016-11-21")
    if err != nil {
        log.Println("failed to send message:", err.Error())
    }

    return nil
}

// ----------------------------------------------------------------------------------
func PitchIn(ctx botmux.Context, args string) (botmux.CommandFunc) {
    err := ctx.SendText(TEXT_PITCHIN)
    if err != nil {
        log.Println("failed to send message:", err.Error())
    }

    return nil
}