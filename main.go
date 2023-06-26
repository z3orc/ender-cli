package main

import (
	"fmt"

	"github.com/z3orc/ender-cli/backup"
	"github.com/z3orc/ender-cli/logger"
)

func main() {
	// root := cli.New(cli.Cli{
	// 	Name:        "ender",
	// 	Description: "A command line tool for managing a Minecraft server.",
	// 	Usage:       "ender <command> [args...]",
	// 	Commands:    commands.All(),
	// 	Debug:       false,
	// })

	// root.Parser()

	backup, err := backup.New()
	if err != nil {
		logger.Error.Println(err)
	}

	fmt.Println(backup)

	// result := console.ParseToLogLine("[17:37:54 INFO]: Building unoptimized datafixer")
	// fmt.Println(result)
}
