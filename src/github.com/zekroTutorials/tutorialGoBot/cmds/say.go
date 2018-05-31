package cmds

import (
	. "strings"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

// Say - Command Funktion für den "say" Command
func Say(bot *discordgo.Session, args []string, channel *discordgo.Channel, guild *discordgo.Guild, author *discordgo.User, message *discordgo.Message) {
	if len(args) < 1 {
		bot.ChannelMessageSendEmbed(channel.ID, &discordgo.MessageEmbed {
			Description: "Invalid arguments",
			Color:       0xf42613,
		})
	}
	
	saycont := Join(args, " ")
	aviurl := fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.png", author.ID, author.Avatar)
	bot.ChannelMessageSendEmbed(channel.ID, &discordgo.MessageEmbed {
		Description: saycont,
		Color:       0x13f4f4,
		Author: &discordgo.MessageEmbedAuthor {
			Name:    author.Username,
			IconURL: aviurl,
		},
	})
}