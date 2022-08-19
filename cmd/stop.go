/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"net"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
	"github.com/z3orc/ender-cli/global"
	"github.com/z3orc/ender-cli/util"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Sends a stop signal to the server",
// 	Long: `A longer description that spans multiple lines and likely contains examples
// and usage of using your command. For example:

// Cobra is a CLI library for Go that empowers applications.
// This application is a tool to generate the needed files
// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		stop()
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func stop(){
	const SockAddr = "./data/ipc.sock"

	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " Stopping server"
	s.FinalMSG = "⛔ Server stopped \n"
	s.Start()

	addr, _ := net.ResolveUnixAddr("unix", SockAddr)
	conn, err := net.DialUnix("unix", nil, addr)
	check(err, "DialUnix")

	conn.Write([]byte("stop\n"))
	conn.Close()

	for {
		if util.IsOpened("127.0.0.1", 25565){
			break
		}
	}

	for {
		if _, err := os.Stat(global.DATA_DIR + "/ipc.sock"); err == nil {
			break
		 }
	}

	s.Stop()
}
