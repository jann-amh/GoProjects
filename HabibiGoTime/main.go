package main

import (
	resources "Documents/Git/GoProjects/HabibiGoTime/JsonController"
	twitch "Documents/Git/GoProjects/HabibiGoTime/Twitch"
)

func main() {
	resources := resources.LoadSettings()
	bot := twitch.NewTwitchBot(resources)
	bot.Connect()
}
