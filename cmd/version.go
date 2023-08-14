package cmd

import (
	"fmt"
	"github.com/hebelsan/go-template-cli/internal/cmd/version"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var (
	printShort = false
	outputType = "json"
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the application version.",
		Long:  `Print the application version.`,
		Run: func(cmd *cobra.Command, args []string) {
			version := version.GenVersionString(printShort, outputType)
			fmt.Print(version)
		},
	}
)

func init() {
	versionCmd.Flags().BoolVarP(&printShort, "short", "s", false, "Print just the version number.")
	versionCmd.Flags().StringVarP(&outputType, "output", "o", "json", "Output format. One of 'yaml' or 'json'.")
	rootCmd.AddCommand(versionCmd)
}
