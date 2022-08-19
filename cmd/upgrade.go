/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/erikgeiser/promptkit/selection"
	"github.com/spf13/cobra"
	"github.com/z3orc/ender-cli/config"
	"github.com/z3orc/ender-cli/endpoints/paper"
	"github.com/z3orc/ender-cli/endpoints/purpur"
	"github.com/z3orc/ender-cli/endpoints/vanilla"
	"github.com/z3orc/ender-cli/global"
	"github.com/z3orc/ender-cli/util"
)

// upgradeCmd represents the upgrade command
var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrades the server to a newer version of Minecraft",
// 	Long: `A longer description that spans multiple lines and likely contains examples
// and usage of using your command. For example:

// Cobra is a CLI library for Go that empowers applications.
// This application is a tool to generate the needed files
// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		upgrade()
	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upgradeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upgradeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func upgrade(){
	current_version := config.Get(global.CONFIG_ENDER_PATH, "version")
	flavour := config.Get(global.CONFIG_ENDER_PATH, "flavour")
	choices := possible_versions(current_version, flavour, false)

	if len(choices) == 0 {
		fmt.Println("🚀  You are already on the latest version!")
		os.Exit(0)
	}

	sp := selection.New("Choose new server version",
	selection.Choices(choices))
	sp.PageSize = 6

	new_version, err := sp.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " Upgrading server \n"
	s.FinalMSG = "✅  Server upgraded! \n"
	spinnerErrorMsg := "Could not upgrade server! \n"
	s.Start()
	time.Sleep(1 * time.Second)

	err = config.Set(global.CONFIG_ENDER_PATH, "version", new_version.String)
	if err != nil {
		s.FinalMSG = spinnerErrorMsg
		s.Stop()
		log.Fatalln(err)
	}

	err = util.DownloadFile(global.JAR_PATH, "https://dynamic.z3orc.com/" + flavour + "/" + new_version.String)
	if err != nil {
		s.FinalMSG = spinnerErrorMsg
		s.Stop()
		config.Set(global.CONFIG_ENDER_PATH, "version", current_version)
		log.Fatalln(err)
	}

	s.Stop()
	
}

func possible_versions(current_version string, flavour string, snapshot bool) []string{
	var possible_versions []string;

	if flavour == "vanilla" {
		versions := vanilla.GetVersions().Versions

		for _, version := range versions {
			if version.Type != "snapshot" {
				possible_versions = append(possible_versions, version.Id)
				if version.Id == current_version {
					break
				}
			}
		}
	} else if flavour == "paper" {
		versions := paper.GetVersions().Versions
		for i, j := 0, len(versions)-1; i < j; i, j = i+1, j-1 {
			versions[i], versions[j] = versions[j], versions[i]
		}

		for index := range versions {
			possible_versions = append(possible_versions, versions[index])
			if versions[index] == current_version {
				break
			}
		}
	} else if flavour == "purpur" {
		versions := purpur.GetVersions().Versions
		for i, j := 0, len(versions)-1; i < j; i, j = i+1, j-1 {
			versions[i], versions[j] = versions[j], versions[i]
		}

		for index := range versions {
			possible_versions = append(possible_versions, versions[index])
			if versions[index] == current_version {
				break
			}
		}
	}

	return possible_versions
}
