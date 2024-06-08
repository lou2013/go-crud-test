package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"test.com/test/internal/app"
)

var rootCmd = &cobra.Command{
	Use:   "help",
	Short: "start the application",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start the application",
	Run: func(cmd *cobra.Command, args []string) {
		app.Start()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(startCmd)
}
