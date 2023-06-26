package console

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/z3orc/ender-cli/logger"
)

var Verbose bool = false

type Console struct {
	cmd    *exec.Cmd
	stdin  *bufio.Writer
	stdout *bufio.Reader
}

func New(cmd *exec.Cmd) *Console {
	c := &Console{
		cmd: cmd,
	}

	stdout, _ := cmd.StdoutPipe()
	c.stdout = bufio.NewReader(stdout)

	stdin, _ := c.cmd.StdinPipe()
	c.stdin = bufio.NewWriter(stdin)

	return c
}

func (c *Console) Write(line string) {
	c.stdin.WriteString(fmt.Sprintf("%s\r\n", line))
}

func (c *Console) Read() (string, error) {
	return c.stdout.ReadString('\n')
}

func (c *Console) Start() {
	logger.Info.Println("Starting server")
	c.cmd.Start()
}

func (c *Console) Stop() {
	logger.Info.Println("Stopping server")
	c.cmd.Process.Signal(os.Interrupt)
	c.cmd.Wait()
	logger.Info.Println("Server stopped")
}
