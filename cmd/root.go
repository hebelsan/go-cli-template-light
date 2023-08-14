package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "my-cli-app",
	Short: "A brief description of your CLI app",
	Long:  `A longer description that explains what your CLI app does`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func GetRootCmd() *cobra.Command {
	return rootCmd
}

func init() {
	// Add global flags and setup
}
