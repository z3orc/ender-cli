package commands

import cli "github.com/z3orc/simple-cli"

var all []*cli.Command

func init() {
	all = []*cli.Command{Start}
}

func All() []*cli.Command {
	return all
}
