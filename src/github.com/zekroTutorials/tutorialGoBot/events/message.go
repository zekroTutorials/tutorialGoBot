package events

import (
	. "strings"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/zekroTutorials/tutorialGoBot/cmds"
	"github.com/zekroTutorials/tutorialGoBot/util"
)


var commands = map[string]func(bot *discordgo.Session, args []string, channel *discordgo.Channel, guild *discordgo.Guild, author *discordgo.User, message *discordgo.Message) {
	"ping": cmds.Ping,
	"say":  cmds.Say,
}


func MessageCreate(bot *discordgo.Session, event *discordgo.MessageCreate) {
	msg := event.Message
	content := msg.Content
	prefix := util.GetConfig().Prefix

	if (HasPrefix(content, prefix)) {
		channel, _ := bot.Channel(msg.ChannelID)
		guild, _ := bot.Guild(channel.GuildID)
		author := msg.Author

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