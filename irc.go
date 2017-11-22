package main

import (
	"fmt"

	irc "github.com/thoj/go-ircevent"
)

var ircObj *irc.Connection

func connectToIRC() {
	ircObj = irc.IRC(Config.IRCInfo.Nick, "d2i")
	ircObj.UseTLS = Config.IRCInfo.IsTLS
	ircObj.AddCallback("001", func(ev *irc.Event) {
		for _, v := range Config.Mapping {
			ircObj.Join(v)
		}
	})
	ircObj.AddCallback("PRIVMSG", ircPrivMsg)
	err := ircObj.Connect(Config.IRCInfo.Addr)
	HandleFatal(err)
}

func ircPrivMsg(ev *irc.Event) {
	go func(event *irc.Event) {
		//event.Message() contains the message
		//event.Nick Contains the sender
		//event.Arguments[0] Contains the channel
		if event.Nick == Config.IRCInfo.Nick {
			return
		}

		cid, ok := Config.RMapping[event.Arguments[0]] // xD
		if !ok {
			return
		}
		DG.ChannelMessageSend(cid, fmt.Sprintf(IRCFormatting, event.Nick, event.Message()))
	}(ev)
}

func ircLoop() {
	ircObj.Loop()
}
