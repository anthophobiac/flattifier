package cmd

import (
	"flattifier/internal/converter"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"path/filepath"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert a JSON file to CSV",
	Run: func(cmd *cobra.Command, args []string) {
		inputFile, _ := cmd.Flags().GetString("input")
		outputFile, _ := cmd.Flags().GetString("output")

		if inputFile == "" || outputFile == "" {
			color.Red("Input and output file paths are required")
			return
		}

		if filepath.Ext(inputFile) != ".json" {
			color.Red("Input file must be a .json file")
			return
		}

		if filepath.Ext(outputFile) != ".csv" {
			color.Red("Output file must be a .csv file")
			return
		}

		if err := converter.ConvertJSONToCSV(inputFile, outputFile); err != nil {
			color.Red("Conversion failed: %v", err)
		} else {
			color.Green("Successfully converted %s to %s", inputFile, outputFile)
		}
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)

	convertCmd.Flags().StringP("input", "i", "", "Input JSON file path")
	convertCmd.Flags().StringP("output", "o", "", "Output CSV file path")
}
