package main

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

var (
	// Config is the global config, sourced from config.toml
	Config = CfgFmt{}
)

// nolint
type DiscordInfo struct {
	BotToken string
}

// nolint
type IRCInfo struct {
	Nick  string
	Addr  string
	IsTLS bool
}

// CfgFmt describes the available formats
type CfgFmt struct {
	DiscordInfo DiscordInfo
	IRCInfo     IRCInfo
	Mapping     map[string]string // Discord ChannelID -> IRC channel
	RMapping    map[string]string
}

func reverseMapping() {
	Config.RMapping = make(map[string]string)
	for k, v := range Config.Mapping {
		Config.RMapping[v] = k // simple asf lmao
	}
}

func init() {
	data, err := ioutil.ReadFile("config.toml")
	HandleFatal(err)
	err = toml.Unmarshal(data, &Config)
	HandleFatal(err)
	reverseMapping()
}
