package jsoncontroller

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadSettings() Resources {
	var resources Resources
	content, err := os.ReadFile("./Resources.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(content, &resources)
	if err != nil {
		panic(err)
	}
	fmt.Println("Loaded resources")
	return resources
}

func LoadCommands(commands *[]string) {
	content, err := os.ReadFile("./Commands.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(content, &commands)
	if err != nil {
		panic(err)
	}
	fmt.Println("Loaded commands")
}
