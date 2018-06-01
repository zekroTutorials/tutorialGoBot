package cmds

import (
	"github.com/bwmarrin/discordgo"
)

// Ping - Command Funktion für den "ping" Command
func Ping(bot *discordgo.Session, args []string, channel *discordgo.Channel, guild *discordgo.Guild, author *discordgo.User, message *discordgo.Message) error {
	embed := discordgo.MessageEmbed {
		Description: ":ping_pong:  Pong!",
		Color: 		 0xb9ea17,
	}
	_, err := bot.ChannelMessageSendEmbed(channel.ID, &embed)
	return err
}