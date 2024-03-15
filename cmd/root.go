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
	RunE: func(cmd *cobra.Command, args []string) error {
		return app.Run(flags)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var flags = app.Options{Sort: app.Name}

func init() {
	rootCmd.Flags().StringSliceVarP(&flags.Locations, "locations", "l", nil, "a comma-separated list of locations to display")
	rootCmd.Flags().VarP(&flags.Sort, "sort", "s", "field to sort the observations by")
	rootCmd.Flags().BoolVarP(&flags.Reverse, "reverse", "r", false, "print the observations in reverse order")
}
