package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"winget/core"
)

var rootCmd = &cobra.Command{
	Use:   "myget",
	Short: "A simple winget wrapper",
}

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import packages from a JSON file",
	Run: func(cmd *cobra.Command, args []string) {
		inputFile, _ := cmd.Flags().GetString("input")
		downloadFlag, _ := cmd.Flags().GetBool("download")

		if inputFile == "" {
			log.Fatal("Input file is required")
		}

		core.Import(inputFile, downloadFlag)
	},
}

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export packages to a JSON file",
	Run: func(cmd *cobra.Command, args []string) {
		exportPath, _ := cmd.Flags().GetString("export")
		downloadFlag, _ := cmd.Flags().GetBool("download")

		if exportPath == "" {
			log.Fatal("Export path is required")
		}

		core.Export(exportPath, downloadFlag)
	},
}

func init() {
	// import command flags
	importCmd.Flags().StringP("input", "i", "./export.json", "Input JSON file containing packages to import")
	importCmd.Flags().BoolP("download", "d", false, "Also download the packages after importing")
	//importCmd.MarkFlagRequired("input")

	// export command flags
	exportCmd.Flags().StringP("export", "e", "./export.json", "Path to export packages to JSON file")
	exportCmd.Flags().BoolP("download", "d", false, "Also download the packages after exporting")

	//exportCmd.MarkFlagRequired("export")

	// Add commands to root
	rootCmd.AddCommand(importCmd)
	rootCmd.AddCommand(exportCmd)

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
