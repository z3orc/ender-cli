package cmd

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/z3orc/ender-cli/config"
	"github.com/z3orc/ender-cli/global"
)

// monitorCmd represents the monitor command
var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Starts the server and opens ipc.",
// 	Long: `A longer description that spans multiple lines and likely contains examples
// and usage of using your command. For example:

// Cobra is a CLI library for Go that empowers applications.
// This application is a tool to generate the needed files
// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		monitor()
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// monitorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// monitorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func monitor(){
	// addr, err := net.ResolveTCPAddr("tcp", "localhost:8000")
	// check(err, "ResolveTCPAddr")

	const SockAddr = "./data/ipc.sock"

	if err := os.RemoveAll(SockAddr); err != nil {
        log.Fatal(err)
    }

	addr, err := net.ResolveUnixAddr("unix", SockAddr)
	check(err, "ResolveUnixAddr")

	conn, err := net.ListenUnix("unix", addr)
	check(err, "ListenUnix")

	fmt.Println("listening on:", conn.Addr())

	ram := config.Get(global.CONFIG_ENDER_PATH, "ram")
	fmt.Println(ram)

	server := exec.Command("java","-Xmx" + ram + "M", "-Xms" + ram + "M", "-jar", "." + global.JAR_PATH, "nogui")
	server.Dir = global.DATA_DIR
	stdout, _ := server.StdoutPipe()
	stdin, _ := server.StdinPipe()

	
	go func() {
		server.Start()
		server.Wait()
		conn.Close()
		defer os.Exit(0)
	}()

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
    // cleanup
    server.Process.Kill()
	conn.Close()
    os.Exit(1)
	}()

	for {
		client, err := conn.AcceptUnix()
		check(err, "AcceptTCP")

		go io.Copy(client, stdout)
		go io.Copy(stdin, client)

		fmt.Println("Accepted from:", client.RemoteAddr())
	}

}

func check(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s error: %v\n", msg, err)
		os.Exit(1)
	}
}
