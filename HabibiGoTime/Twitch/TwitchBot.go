package twitch

import (
	settings "Documents/Git/GoProjects/HabibiGoTime/JsonController"
	"strings"

	irc "github.com/gempir/go-twitch-irc/v2"

	"fmt"
)

const (
	Prefix = "go"
)

type TwitchBot struct {
	twitchClient irc.Client
	Resources    settings.Resources
	Commands     *[]string
}

func NewTwitchBot(resources settings.Resources, cmds *[]string) *TwitchBot {
	twitchBot := &TwitchBot{
		twitchClient: *irc.NewClient(resources.Username, resources.OAuthToken),
		Resources:    resources,
		Commands:     cmds,
	}

	twitchBot.twitchClient.OnConnect(func() {
		fmt.Println("Bot Connected")
	})

	twitchBot.twitchClient.OnPrivateMessage(func(message irc.PrivateMessage) {
		fmt.Println(message.Channel, "|", message.User.Name+":", message.Message)
		twitchBot.HandleCommand(message)
	})

	return twitchBot
}

func (twitchBot *TwitchBot) Connect() {
	for _, c := range twitchBot.Resources.Channels {
		twitchBot.twitchClient.Join(c)
	}
	twitchBot.twitchClient.Connect()
}

func (twitchBot *TwitchBot) Send(channel string, message string) {
	twitchBot.twitchClient.Say(channel, message)
}

func (twitchBot *TwitchBot) HandleCommand(message irc.PrivateMessage) {
	cmds := *twitchBot.Commands
	if strings.HasPrefix(message.Message, Prefix+cmds[0]) {
		twitchBot.Send(message.Channel, "gopherDance")
	}
}
