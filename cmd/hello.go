package cmd

import (
	"cobra-template/pkg/logger"
	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Prints a hello message",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("Hello from the cobra cli app!")
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
