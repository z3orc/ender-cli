package cmd

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/erikgeiser/promptkit/selection"
	"github.com/erikgeiser/promptkit/textinput"
	"github.com/spf13/cobra"
	"github.com/z3orc/ender-cli/global"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		setup()
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func setup(){
	dir_res := create_directories()
	if dir_res != 0 {
		log.Fatal("Could not create directories")
	}

	// dl_res  := download_file(global.JAR_PATH, "http://dynamic.z3orc.com/paper/1.19")
	// if dl_res != nil {
	// 	log.Fatal("Could not download jar file")
	// }

	sp := selection.New("Which server flavour?",
	selection.Choices([]string{"Vanilla", "Paper", "Purpur"}))
	sp.PageSize = 3

	choice, err := sp.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	// do something with the final choice
	_ = choice


	input := textinput.New("Which minecraft version?")
	input.InitialValue = "1.19"
	input.Placeholder = "Version cannot be empty"

	name, err := input.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	// do something with the result
	_ = name


	input = textinput.New("How much ram should the server use (in Mb, do not include the unit)")
	input.InitialValue = "2000"
	input.Placeholder = "Ram cannot be empty"

	name, err = input.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	// do something with the result
	_ = name


	input = textinput.New("Enter the maximum player limit")
	input.InitialValue = "12"
	input.Placeholder = "Player limit cannot be empty"

	name, err = input.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	// do something with the result
	_ = name
}

func create_directories() int{

	dirs := [4]string{global.BACKUP_DIR, global.BIN_DIR, global.CONFIG_DIR, global.DATA_DIR}

	for _, dir := range dirs {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return 1
		}
	}
	return 0

}

func download_file(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	status := resp.StatusCode
	if status != 200 {
		return errors.New("could not download jar file")
	}

	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
