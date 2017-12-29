package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// nolint
var DG *discordgo.Session

// IRCmaxLen represents the maximum length of a IRC message
const IRCmaxLen = 256

// IRCsplit splits a message over 480 bytes to chunks of 480-byte messages.
func IRCsplit(s string) [][]byte {
	var r [][]byte
	b := []byte(s)
	for len(b) > 0 {
		l := len(b)
		if l > IRCmaxLen {
			l = IRCmaxLen
		}
		r = append(r, b[:l:l])
		b = b[l:]
	}
	return r
}

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
	// revive autism
	if strings.Contains(msg, "\r\n") {
		chunks = strings.Split(msg, "\r\n")
	} else {
		chunks = []string{msg}
	}

	// not managing the character limit per messages, so we do that now
	chunkcp := []string{}

	for _, chunk := range chunks {
		for _, chk1 := range IRCsplit(chunk) {
			chunkcp = append(chunkcp, string(chk1))
		}
	}

	for _, chunk := range chunkcp {
		ircObj.Privmsg(im, chunk)
	}
}

func closeDiscord() {
	DG.Close()
}
