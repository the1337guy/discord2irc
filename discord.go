package main

import (
	"fmt"
	"unicode/utf8"

	"github.com/bwmarrin/discordgo"
)

// nolint
var DG *discordgo.Session

func createDiscordSession() {
	var err error
	DG, err = discordgo.New("Bot " + Config.DiscordInfo.BotToken)
	HandleFatal(err)
	DG.AddHandler(discordOnMessage)
	err = DG.Open()
	HandleFatal(err)
}

func discordOnMessage(dg *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == dg.State.User.ID {
		return
	}

	im, ok := Config.Mapping[m.ChannelID]
	msg := m.ContentWithMentionsReplaced()

	if !ok {
		return
	}

	msg = fmt.Sprintf(DiscordFormatting, m.Author.String(), msg)

	var chunks []string
	for len(msg) > 512 {
		i := 512
		for i >= 512-utf8.UTFMax && !utf8.RuneStart(msg[i]) {
			i--
		}
		chunks = append(chunks, msg[:i])
		msg = msg[i:]
	}
	if len(msg) > 0 {
		chunks = append(chunks, msg) // nolint
	}

	for _, chunk := range chunks {
		ircObj.Privmsg(im, chunk)
	}
}

func closeDiscord() {
	DG.Close()
}
