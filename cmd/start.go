/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
	"github.com/z3orc/ender-cli/util"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the server",
// 	Long: `A longer description that spans multiple lines and likely contains examples
// and usage of using your command. For example:

// Cobra is a CLI library for Go that empowers applications.
// This application is a tool to generate the needed files
// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func start(){
	executable, err := os.Executable();
	if err != nil {
		fmt.Println("Error: ", err)
	}

	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " Starting server"
	s.FinalMSG = "🚀 Server started \n"
	s.Start()

	if runtime.GOOS == "windows" {
		monitor := exec.Command(executable, "monitor")
		monitor.Stdout = os.Stdout
		monitor.Stdin = os.Stdin
		monitor.Stderr = os.Stderr
		monitor.Run()
	} else {
		monitor := exec.Command(executable, "monitor")
		monitor.Start()
	}

	for {
		if util.IsOpened("127.0.0.1", 25565){
			break
		}
	}

	s.Stop()
}
