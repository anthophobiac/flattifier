package cmd

import (
	"fmt"
	"os"

	"cobra-template/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cobra-template",
	Short: "A starter template for Go CLI tools using Cobra",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("config", "", "config file (default is .env)")
}
