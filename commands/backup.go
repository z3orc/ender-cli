package commands

import (
	"github.com/z3orc/ender-cli/backup"
	"github.com/z3orc/ender-cli/logger"
	cli "github.com/z3orc/simple-cli"
)

var Backup *cli.Command = &cli.Command{
	Name:        "backup",
	Usage:       "backup",
	Description: "Makes a copy of all the server files",
	Arguments:   0,
	Options: []*cli.Option{
		{
			Name:        "purge",
			ShortName:   "p",
			Description: "Removes all backups",
			Arguments:   1,
		},
	},
	Run: func(s [][]string) {
		if len(s) > 0 && (s[0][0] == "--purge" || s[0][0] == "-p") {
			logger.Info.Println("Purging backups")
			if err := backup.PurgeBackups(); err != nil {
				logger.Error.Fatalln("could not purge all backups")
			} else {
				logger.Info.Println("Backups purged")
			}
			return
		}
		logger.Info.Println("Creating new backup")
		_, err := backup.New()
		if err != nil {
			logger.Error.Fatalln("could not create backup. " + err.Error())
		} else {
			logger.Info.Println("New backup created")
		}
	},
}
