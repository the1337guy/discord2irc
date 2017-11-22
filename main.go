package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// nolint
var IRCFormatting = "<**%s**> %s"

// nolint
var DiscordFormatting = "<\x02%s\x02>: %s"

func main() {
	createDiscordSession()
	connectToIRC()

	go ircLoop()

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	closeDiscord()
}
