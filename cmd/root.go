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
	Long: `Shows the current UV index for various locations around Australia.
Observations are sourced from ARPANSA.
See disclaimer: https://www.arpansa.gov.au/our-services/monitoring/ultraviolet-radiation-monitoring/ultraviolet-radation-data-information#Disclaimer`,
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
	rootCmd.Flags().StringSliceVarP(&flags.Locations, "locations", "l", nil, "comma-separated list of locations to display (accepts id and name)")
	rootCmd.Flags().VarP(&flags.Sort, "sort", "s", "field to sort observations by")
	rootCmd.Flags().BoolVarP(&flags.Reverse, "reverse", "r", false, "print observations in reverse order")
}
