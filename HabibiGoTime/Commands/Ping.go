package commands

import (
	twitchBot "Documents/Git/GoProjects/HabibiGoTime/Twitch"

	irc "github.com/gempir/go-twitch-irc/v2"
)

func HandlePing(bot *twitchBot.TwitchBot, chatMessage irc.PrivateMessage) {
	bot.Send(chatMessage.Channel, "Pong")
}
