package console

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/z3orc/ender-cli/logger"
)

var Verbose bool = false

type Wrapper struct {
	cmd    *exec.Cmd
	stdin  *bufio.Writer
	stdout *bufio.Reader
}

func New(cmd *exec.Cmd) *Wrapper {
	c := &Wrapper{
		cmd: cmd,
	}

	stdout, _ := cmd.StdoutPipe()
	c.stdout = bufio.NewReader(stdout)

	stdin, _ := c.cmd.StdinPipe()
	c.stdin = bufio.NewWriter(stdin)

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
}

func (c *Wrapper) Stop() {
	logger.Info.Println("Stopping server")
	c.cmd.Process.Signal(os.Interrupt)
	c.cmd.Wait()
	logger.Info.Println("Server stopped")
}

func (c *Wrapper) Wait() {
	c.cmd.Wait()
}
