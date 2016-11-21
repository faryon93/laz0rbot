package ikr

import (
	"log"
	"strings"

	"github.com/faryon93/laz0rbot/botmux"

    //"gopkg.in/telegram-bot-api.v4"
)


// ----------------------------------------------------------------------------------
//  telegram commands
// ----------------------------------------------------------------------------------

func Entry(ctx botmux.Context, args string) (botmux.CommandFunc) {
	switch strings.ToLower(args) {
		case "schedule":
			log.Println("view schedule")
	}

	return nil
}