package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/YoshinoriN/cahsper-cli/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCommand.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize cahsper-cli",
	Long:  "initialize cahsper-cli",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {

		homeDir := utils.GetUserHomeDirectory()
		configDir := filepath.Join(homeDir, ".cahsper")
		configFilePath := filepath.Join(configDir, ".config")

		if !utils.Exists(configDir) {
			utils.MakeDir(configDir)
		}

		if !utils.Exists(configFilePath) {
			file, err := os.OpenFile(configFilePath, os.O_WRONLY|os.O_CREATE, 0600)
			if err != nil {
				panic(err)
			}
			fmt.Println("created config file: " + file.Name())
		} else {
			fmt.Println("config file is already exists: " + configFilePath)
		}
		return nil
	},
}
