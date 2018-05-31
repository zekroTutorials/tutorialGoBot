package cmds

import (
	"github.com/bwmarrin/discordgo"
)

func Ping(bot *discordgo.Session, args []string, channel *discordgo.Channel, guild *discordgo.Guild, author *discordgo.User, message *discordgo.Message) {
	embed := discordgo.MessageEmbed {
		Description: ":ping_pong:  Pong!",
		Color: 		 0xb9ea17,
	}
	bot.ChannelMessageSendEmbed(channel.ID, &embed)
}