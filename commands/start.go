package commands

import (
	"os"
	"os/exec"
	"os/signal"
	"time"

	"github.com/z3orc/ender-cli/backup"
	"github.com/z3orc/ender-cli/logger"
	wrapper "github.com/z3orc/ender-cli/wrapper"
	cli "github.com/z3orc/simple-cli"
)

var Start *cli.Command = &cli.Command{
	Name:        "start",
	Usage:       "start",
	Description: "Starts up the minecraft server with current configuration",
	Arguments:   0,
	Options: []*cli.Option{
		{
			Name:        "verbose",
			ShortName:   "v",
			Description: "Write out every log entry to stdout, useful for debugging a non-functioning server",
			Arguments:   0,
		},
	},
	Run: func(s [][]string) {
		if len(s) > 0 && (s[0][0] == "--verbose" || s[0][0] == "-v") {
			wrapper.Verbose = true
		}
		start()
	},
}

func start() {
	javaExec := exec.Command("java", "-jar", "server.jar", "nogui")
	javaExec.Dir = "./testing"
	javaExec.Stdin = os.Stdin

	server := wrapper.New(javaExec)
	server.Start()
	go server.ReadLogs()

	quitSignal := make(chan os.Signal, 1)
	signal.Notify(quitSignal, os.Interrupt)

	stoppedSignal := make(chan int)
	go func() {
		server.Wait()
		stoppedSignal <- 1
	}()

	backupSignal := make(chan int)
	go func() {
		time.Sleep(24 * time.Hour)
		backupSignal <- 1
	}()

	select {
	case <-stoppedSignal:
		os.Exit(0)
	case <-quitSignal:
		server.Stop()
	case <-backupSignal:
		server.Stop()
		_, err := backup.New()
		if err != nil {
			logger.Error.Fatalln(err.Error())
		} else {
			logger.Info.Println("New backup created")
		}

		start()
	}
}
