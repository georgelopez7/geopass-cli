package cmd

import (
	"fmt"
	"geopass-cli/config"
	"geopass-cli/generator"

	"github.com/spf13/cobra"
)

var (
	length int
)

const MinPasswordLength = config.DefaultPasswordLength

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "generates a secure password",
	Run: func(cmd *cobra.Command, args []string) {

		if length < MinPasswordLength {
			msg := fmt.Sprintf("Length must be at least %d.", MinPasswordLength)
			fmt.Println("Error:", msg)
			return
		}

		password, entropy, err := generator.GeneratePassword(length)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("\nGenerated password:" + config.Cyan + "\n" + password + config.Reset)
		entropyStr := fmt.Sprintf("%.2f", entropy)
		fmt.Println("\nPassword entropy:" + config.Green + "\n" + entropyStr + config.Reset)
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.Flags().IntVarP(&length, "length", "l", MinPasswordLength, "length of the password")
}
