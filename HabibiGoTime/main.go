package main

import (
	jsoncontroller "Documents/Git/GoProjects/HabibiGoTime/JsonController"
	twitch "Documents/Git/GoProjects/HabibiGoTime/Twitch"
)

func main() {
	var commands []string
	resources := jsoncontroller.LoadSettings()
	jsoncontroller.LoadCommands(&commands)
	bot := twitch.NewTwitchBot(resources, &commands)
	bot.Connect()
}
