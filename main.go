package main

import (
	"github.com/HaseemKhattak01/mydriveuploader/cmd"
	"github.com/HaseemKhattak01/mydriveuploader/config"
)

func main() {
	config.LoadConfig()
	cmd.Execute()
}
