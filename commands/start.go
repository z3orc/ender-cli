package commands

import (
	"bufio"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strings"

	"github.com/z3orc/ender-cli/console"
	"github.com/z3orc/ender-cli/logger"
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
			console.Verbose = true
		}
		start()
	},
}

func start() {
	javaExec := exec.Command("java", "-jar", "server.jar", "nogui")
	javaExec.Dir = "./testing"
	javaExec.Stdin = os.Stdin
	// server.Stdout = os.Stdout
	// server.Start()

	server := console.New(javaExec)
	server.Start()

	go func() {
		for {
			line, err := server.Read()
			if err != nil {
				if err == io.EOF {
					return
				} else {
					logger.Error.Println("could not read console output")
				}
			}

			console.Parse(line)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		scanner.Scan()
		input := scanner.Text()
		input = strings.Trim(input, " ")

		if input == "stop" {
			server.Stop()
			os.Exit(0)
		} else if input != "" {
			server.Write(input + "\n")
		}

	}()

	backup := make(chan int)
	// go func() {
	// 	time.Sleep(60 * time.Second)
	// 	backup <- 1
	// }()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	select {
	case <-quit:
		server.Stop()
	case <-backup:
		server.Stop()
		start()
	}
}
