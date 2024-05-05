package main

import "github.com/timohahaa/userbot/internal/app"

const configFilePath = "./config/config.yaml"

func main() {
	app.Run(configFilePath)
}
