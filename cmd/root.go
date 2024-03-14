package cmd

import (
	"fmt"
	"github.com/msharris/uv/app"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "uv",
	Short: "Shows the current UV index for various locations around Australia",
	Long: `uv is an app that shows the current UV index for various locations around Australia.
UV observations are sourced from the ARPANSA API.`,
	Run: func(cmd *cobra.Command, args []string) {
		app.Run()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Flags will go here
}
