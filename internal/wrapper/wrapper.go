package wrapper

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/z3orc/ender-cli/internal/logger"
)

var Verbose bool = false

type Wrapper struct {
	cmd     *exec.Cmd
	stdin   *bufio.Writer
	stdout  *bufio.Reader
	Stopped chan int
}

func New(cmd *exec.Cmd) *Wrapper {
	c := &Wrapper{
		cmd: cmd,
	}

	stdout, _ := cmd.StdoutPipe()
	c.stdout = bufio.NewReader(stdout)

	stdin, _ := c.cmd.StdinPipe()
	c.stdin = bufio.NewWriter(stdin)

	c.Stopped = make(chan int)

	return c
}

func (c *Wrapper) Write(line string) {
	c.stdin.WriteString(fmt.Sprintf("%s\r\n", line))
	c.stdin.Flush()
}

func (c *Wrapper) Read() (string, error) {
	return c.stdout.ReadString('\n')
}

func (c *Wrapper) Start() {
	logger.Info.Println("Starting server")
	c.cmd.Start()
	stoppedSignal := make(chan int)
	go func() {
		c.Wait()
		stoppedSignal <- 1
	}()
}

func (c *Wrapper) Stop() {
	logger.Info.Println("Stopping server")
	c.cmd.Process.Signal(os.Interrupt)
	c.Wait()
	logger.Info.Println("Server stopped")
}

func (c *Wrapper) Wait() {
	c.cmd.Wait()
}

func (c *Wrapper) ReadLogs() {
	for {
		line, err := c.Read()
		if err != nil {
			if err != io.EOF {
				logger.Error.Println("could not read console output")
			}
			return
		}
		Parse(line)
	}
}
