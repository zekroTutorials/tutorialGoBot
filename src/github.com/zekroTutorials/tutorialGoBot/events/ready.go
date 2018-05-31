package events

import (
	. "fmt"
	"github.com/bwmarrin/discordgo"
)

func Ready(bot *discordgo.Session, event *discordgo.Ready) {
	user := event.User
	Printf("Connected as %s#%s (%s)\n", user.Username, user.Discriminator, user.ID)
	bot.UpdateStatus(0, "GO <3")
}