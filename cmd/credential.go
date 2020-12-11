package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/YoshinoriN/cahsper-cli/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCommand.AddCommand(credentialCommand)
	credentialCommand.AddCommand(showCredentialCommmand)
	credentialCommand.AddCommand(setCredentialCommand)
}

var credentialCommand = &cobra.Command{
	Use:   "credential",
	Short: "get/set your credential for cahsper",
	Long:  "get/set your credential for cahsper",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// TODO: refactor
var showCredentialCommmand = &cobra.Command{
	Use:   "list",
	Short: "Show all of cahsper credential variables.",
	Long:  "Show all of cahsper credential variables.",
	Run: func(cmd *cobra.Command, args []string) {

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("UserName: ")
		scanner.Scan()
		userName := scanner.Text()
		if userName == "" {
			fmt.Print("UserName required.")
			os.Exit(0)
		}

		var err error
		var password = ""
		userName, password, err = utils.GetAccount(userName)
		if err != nil {
			if strings.Contains(fmt.Sprintln(err), "secret not found in keyring") {
				fmt.Printf("UserName %s does not exists.\n", userName)
				os.Exit(0)
			} else {
				log.Fatal(err)
				os.Exit(1)
			}
		}
		fmt.Println("Password: ", password)

		idToken, err := utils.GetCredential(userName, utils.IDToken)
		if err != nil {
			if !strings.Contains(fmt.Sprintln(err), "secret not found in keyring") {
				log.Fatal(err)
				os.Exit(1)
			}
		}
		fmt.Println("IdToken: ", idToken)

		accessToken, err := utils.GetCredential(userName, utils.AccessToken)
		if err != nil {
			if !strings.Contains(fmt.Sprintln(err), "secret not found in keyring") {
				log.Fatal(err)
				os.Exit(1)
			}
		}
		fmt.Println("AccessToken: ", accessToken)

		refreshToken, err := utils.GetCredential(userName, utils.RefreshToken)
		if err != nil {
			if !strings.Contains(fmt.Sprintln(err), "secret not found in keyring") {
				log.Fatal(err)
				os.Exit(1)
			}
		}
		fmt.Println("RefreshToken: ", refreshToken)
	},
}

// TODO: refactor
var setCredentialCommand = &cobra.Command{
	Use:   "set",
	Short: "Set cahsper credential variables.",
	Long:  `Set cahsper credential variables.`,
	Run: func(cmd *cobra.Command, args []string) {

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("UserName: ")
		scanner.Scan()
		userName := scanner.Text()
		if userName == "" {
			fmt.Print("UserName required.")
			os.Exit(0)
		}

		var err error
		var password = ""
		userName, password, err = utils.GetAccount(userName)
		if err != nil {
			if strings.Contains(fmt.Sprintln(err), "secret not found in keyring") {
				fmt.Printf("UserName %s does not exists. Continue to creating new account by %s.\n", userName, userName)
			} else {
				log.Fatal(err)
				os.Exit(1)
			}
		}
		utils.InteractInputHelper("Password", utils.Account, userName, password)
	},
}
