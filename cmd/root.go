package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "1.0.0"

var rootCmd = &cobra.Command{
	Use:   "geopass",
	Short: "A CLI tool for generating secure passwords!",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().BoolP("version", "v", false, "prints the version of geopass.")
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		v, _ := cmd.Flags().GetBool("version")
		if v {
			fmt.Printf("geopass version %s\n", version)
		} else {
			_ = cmd.Help()
		}
	}
}
