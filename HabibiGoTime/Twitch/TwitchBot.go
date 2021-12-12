package Twitch

import (
	settings "Documents/Git/GoProjects/HabibiGoTime/JsonController"

	irc "github.com/gempir/go-twitch-irc/v2"

	"fmt"
)

type TwitchBot struct {
	twitchClient irc.Client
	resources    settings.Resources
}

func NewTwitchBot(resources settings.Resources) TwitchBot {
	twitchBot := TwitchBot{
		twitchClient: *irc.NewClient(resources.Username, resources.OAuthToken),
		resources:    resources,
	}

	twitchBot.twitchClient.OnConnect(func() {
		fmt.Println("Bot Connected")
	})

	twitchBot.twitchClient.OnPrivateMessage(func(message irc.PrivateMessage) {
		fmt.Println(message.Channel, "|", message.User.Name+":", message.Message)
	})
	return twitchBot
}

func (twitchBot TwitchBot) Connect() {
	for _, c := range twitchBot.resources.Channels {
		twitchBot.twitchClient.Join(c)
	}
	twitchBot.twitchClient.Connect()
}
