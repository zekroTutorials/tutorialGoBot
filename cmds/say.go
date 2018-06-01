package cmds

import (
	. "strings"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

// Say - Command Funktion für den "say" Command
func Say(bot *discordgo.Session, args []string, channel *discordgo.Channel, guild *discordgo.Guild, author *discordgo.User, message *discordgo.Message) error {
	var err error
	// Wenn keine Argumente übergeben wurden wird dies als Error Message ausgegeben
	// und die Funktion danach verlassen.
	if len(args) < 1 {
		_, err = bot.ChannelMessageSendEmbed(channel.ID, &discordgo.MessageEmbed {
			Description: "Invalid arguments",
			Color:       0xf42613,
		})
	}
	
	// Alle Argumente werden wieder zu einem String zusammen gefügt (re-assembled)
	saycont := Join(args, " ")
	// Die Avatar-Url muss zuerst aus der User ID und dem Avatar Hash zusammen gestellt werden
	aviurl := fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.png", author.ID, author.Avatar)
	// Danach wird alles als Embed Message ausgegeben.
	_, err = bot.ChannelMessageSendEmbed(channel.ID, &discordgo.MessageEmbed {
		Description: saycont,
		Color:       0x13f4f4,
		Author: &discordgo.MessageEmbedAuthor {
			Name:    author.Username,
			IconURL: aviurl,
		},
	})
	return err
}