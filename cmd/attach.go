package cmd

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/spf13/cobra"
)

// attachCmd represents the attach command
var attachCmd = &cobra.Command{
	Use:   "attach",
	Short: "Attaches to the server console",
// 	Long: `A longer description that spans multiple lines and likely contains examples
// and usage of using your command. For example:

// Cobra is a CLI library for Go that empowers applications.
// This application is a tool to generate the needed files
// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		attach()
	},
}

func init() {
	rootCmd.AddCommand(attachCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// attachCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// attachCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func attach(){
	// addr, _ := net.ResolveTCPAddr("tcp", "localhost:8000")
	// conn, err := net.DialTCP("tcp", nil, addr)
	// check(err, "DialTCP")

	const SockAddr = "./data/ipc.sock"

	addr, _ := net.ResolveUnixAddr("unix", SockAddr)
	conn, err := net.DialUnix("unix", nil, addr)
	check(err, "DialUnix")

	defer conn.Close()

	var data [512]byte

	go func ()  {
		for {
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			input := scanner.Text()
			conn.Write([]byte(input + "\n"))
		}
	}()

	for {
		n, err := conn.Read(data[:])

		if err != nil {
			if err != io.EOF {
				fmt.Println("client read error:", err)
			}
			break
		}

		fmt.Print(string(data[:n]))
	}
}