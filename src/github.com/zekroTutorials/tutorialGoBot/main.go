package main

import (
	"os"
	"fmt"
	"syscall"
	"os/signal"
	"github.com/zekroTutorials/tutorialGoBot/util"
	"github.com/zekroTutorials/tutorialGoBot/events"
	"github.com/bwmarrin/discordgo"
)


func check(err error) {
	if err != nil {
		panic(err)
	}
}

func addHandlers(bot *discordgo.Session) {
	bot.AddHandler(events.Ready)
	bot.AddHandler(events.MessageCreate)
}


func main() {

	util.LoadConfig()
	config := util.GetConfig()

	bot, err := discordgo.New("Bot " + config.Token)
	check(err)

	addHandlers(bot)
	
	fmt.Println("Connecting...")
	err = bot.Open()
	check(err)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	bot.Close()

}