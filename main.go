package main

import (
	"os"
	"fmt"
	"syscall"
	"os/signal"
	"github.com/bwmarrin/discordgo"
	// internal packages
	"github.com/zekroTutorials/tutorialGoBot/util"
	"github.com/zekroTutorials/tutorialGoBot/events"
)

// check dienst intern zur Prüfung von
// errors.
// panic gibt den error stack in der Konsole aus
// und das programm wird gestoppt.
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// addHandler added angegebene Event Handler
// zu der übergebenen Bot Session.
func addHandlers(bot *discordgo.Session) {
	bot.AddHandler(events.Ready)
	bot.AddHandler(events.MessageCreate)
}

func main() {

	// Config laden und die Instanz davon lokal
	// als Pointer speichern
	// Das hat den Vorteil, dass die Konfig nicht
	// kopiert werden muss und im Zweifel bearbeitet
	// werden kann.
	config := util.LoadConfig()

	// Bot Session wird instanziert mit Token aus
	// der Config.
	bot, err := discordgo.New("Bot " + config.Token)
	check(err)

	// Handler hinzufügen
	addHandlers(bot)
	
	fmt.Println("Connecting...")
	// Der bot connected sich zur Discord API
	// mit den angegebenen Daten (token)
	err = bot.Open()
	check(err)

	// An der stelle wird ein Channel mit der Pufferlänge 1 geöffnet,
	// in welchen ein Signal gespeichert wird, wenn der Prozess beendet wird.
	// Danach wird der Channel ausgelesen und die Connection geschlossen.
	// Solange der Channel auf den Eingang des Signals wartet bleibt der Thread
	// bestehen. Ohne dies würde das Programm sonst direkt schließen und es
	// gäbe keinen Event-Loop.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	bot.Close()

}