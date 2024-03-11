/*
Copyright © 2024 Iván Recasens Lahoz <irecasensl@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Acces and modify configuration",
	Long:  `Provides acces to internal variables to customize command behaviour, like file directory`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(configCmd)

}
