package main

import (
	"github.com/z3orc/ender-cli/commands"
	cli "github.com/z3orc/simple-cli"
)

func main() {
	root := cli.New(cli.Cli{
		Name:        "ender",
		Description: "A command line tool for managing a Minecraft server.",
		Usage:       "ender <command> [args...]",
		Commands:    commands.All(),
		Debug:       false,
	})

	root.Parser()
}
