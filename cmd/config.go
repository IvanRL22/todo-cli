/*
Copyright © 2024 Iván Recasens Lahoz <irecasensl@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var verbose bool

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Acces and modify configuration",
	Long: `Provides acces to internal variables to customize command behaviour,
	like file directory`,
	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
			fmt.Print(GetConfigFileVerbose())
		} else {
			fmt.Printf("WIP\nProvides acces to configuration functionality\n")
		}

	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}
