package main

import (
	"mydriveuploader/cmd"
	"mydriveuploader/config"
)

func main() {
	config.LoadConfig()
	cmd.Execute()
}