package events

import (
	. "strings"
	"fmt"
	"github.com/bwmarrin/discordgo"
	// internal imports
	"github.com/zekroTutorials/tutorialGoBot/cmds"
	"github.com/zekroTutorials/tutorialGoBot/util"
)

// Typdefinition CmdFunc für die Funktionen der Commands
type CmdFunc func(bot *discordgo.Session, 
				  args []string, 
				  channel *discordgo.Channel, 
				  guild *discordgo.Guild, 
				  author *discordgo.User, 
				  message *discordgo.Message)

// Commands map in welcher alle Funktionen der Commands zu den zugehörigen
// invokes aufgelistet sind.
var commands = map[string]CmdFunc {
	"ping": cmds.Ping,
	"say":  cmds.Say,
}

// MessgaeCreate Event Handler - wird aufgerufen, wenn eine Message gesendet wurde
func MessageCreate(bot *discordgo.Session, event *discordgo.MessageCreate) {
	msg := event.Message
	content := msg.Content
	prefix := util.GetConfig().Prefix

	// Erst wird geprüft, ob die Message mit dem in der Config festgelegten
	// Prefix beginnt.
	if (HasPrefix(content, prefix)) {
		channel, _ := bot.Channel(msg.ChannelID)
		guild, _ := bot.Guild(channel.GuildID)
		author := msg.Author

		// Ist der Author der Message ein Bot oder wurde die Message nicht
		// in einem normalen Guild Text Channel versendet, so wird an dieser
		// stelle die Funktion verlassen.
		if author.Bot || channel.Type != 0 {
			return
		}

		contsplit := Split(content, " ")
		invoke := contsplit[0][len(prefix):]
		args := contsplit[1:]

		if cmdfunc, ok := commands[invoke]; ok {
			cmdfunc(bot, args, channel, guild, author, msg)
			fmt.Printf("CMD | %s#%s (%s) @ %s (%s) | '%s'",
				author.Username, author.Discriminator, author.ID, guild.Name, guild.ID, content)
		}
	}
}