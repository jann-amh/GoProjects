package commands

import (
	twitchBot "Documents/Git/GoProjects/HabibiGoTime/Twitch"

	irc "github.com/gempir/go-twitch-irc/v2"
)

func Handle(bot *twitchBot.TwitchBot, chatMessage irc.PrivateMessage) {
	if chatMessage.Message == "goping" {

	}
}
