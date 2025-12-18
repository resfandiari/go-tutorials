package cmd

import (
	"fmt"
	"godupe/internal/scanner"
	"godupe/internal/utils"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "godupe",
	Short: "godupe is a tool to find duplicate files in a directory",
	Long:  `godupe scans a specified directory and identifies duplicate files based on their content.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dir := args[0]
		// Implementation of the command's functionality goes here
		color.Cyan("🔍 Starting scan in directory: %s", dir)
		duplicates, err := scanner.FindDuplicates(dir)

		if err != nil {
			color.Red("❌ Error scanning directory: %v", err)
			os.Exit(1)
		}

		if len(duplicates) == 0 {
			color.Green("✅ Great! No duplicate files found.")
			return
		}

		utils.PrintResult(duplicates)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
