package main

import (
	"fmt"

	"github.com/z3orc/ender-cli/internal/backup"
	"github.com/z3orc/ender-cli/internal/global"
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

	br, err := backup.LoadRegistry(global.WORK_DIR + "/backups/backups.json")
	if err != nil {
		fmt.Println("well this did not work... " + err.Error())
	}

	fmt.Println(br)
}
