package cmd

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/erikgeiser/promptkit/confirmation"
	"github.com/erikgeiser/promptkit/selection"
	"github.com/erikgeiser/promptkit/textinput"
	"github.com/spf13/cobra"
	"github.com/z3orc/ender-cli/config"
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

	new_config := make(map[string]string)

	dir_res := create_directories()
	if dir_res != 0 {
		log.Fatal("Could not create directories")
	}

	sp := selection.New("Which server flavour?",
	selection.Choices([]string{"Vanilla", "Paper", "Purpur"}))
	sp.PageSize = 3

	choice, err := sp.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	// do something with the final choice
	new_config["flavour"] = strings.ToLower(choice.String)


	input := textinput.New("Which minecraft version?")
	input.InitialValue = "1.19"
	input.Placeholder = "Version cannot be empty"

	version, err := input.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	// do something with the result
	new_config["version"] = version


	input = textinput.New("How much ram should the server use (in Mb, do not include the unit)")
	input.InitialValue = "2000"
	input.Placeholder = "Ram cannot be empty"

	ram, err := input.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	// do something with the result
	new_config["ram"] = ram


	input = textinput.New("Enter the maximum player limit")
	input.InitialValue = "12"
	input.Placeholder = "Player limit cannot be empty"

	limit, err := input.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	// do something with the result
	new_config["limit"] = limit


	input = textinput.New("Enter a world seed (optional)")
	input.Validate = func (s string) error  {
		return nil
	}
	input.InitialValue = ""

	seed, err := input.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	// do something with the result
	new_config["world_seed"] = seed


	sp = selection.New("Choose default gamemode",
	selection.Choices([]string{"survival", "creative", "adventure"}))
	sp.PageSize = 3

	choice, err = sp.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	// do something with the final choice
	new_config["gamemode"] = choice.String


	sp = selection.New("Choose server difficulty",
	selection.Choices([]string{"peaceful", "easy", "normal", "hard"}))
	sp.PageSize = 3

	choice, err = sp.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	// do something with the final choice
	new_config["difficulty"] = choice.String


	input = textinput.New("Which port the server should listen on")
	input.InitialValue = "25565"

	port, err := input.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	// do something with the result
	new_config["port"] = port


	conf := confirmation.New("Should the server be whitelisted?", confirmation.Undecided)

	whitelist, err := conf.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	// do something with the result
	if whitelist {
		new_config["whitelist"] = "true"
	} else {
		new_config["whitelist"] = "false"
	}


	conf = confirmation.New("I have read and agree to the Minecraft EULA (https://www.minecraft.net/eula)", confirmation.Undecided)

	eula, err := conf.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	// do something with the result
	if eula {
		new_config["eula"] = "true"
	} else {
		new_config["eula"] = "false"
	}

	// err = config.Create(global.CONFIG_ENDER_PATH, config)
	err = config.Create(global.CONFIG_ENDER_PATH, new_config)
	if err != nil {
		log.Fatal(err)
	}

	download_url := fmt.Sprint("http://dynamic.z3orc.com/", new_config["flavour"], "/", new_config["version"])
	err = download_file(global.JAR_PATH, download_url)
	if err != nil {
		log.Fatal(err)
	}
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
	fmt.Println(status)
	if !(status == 200 || status == 302) {
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
