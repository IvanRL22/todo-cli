/*
Copyright © 2024 Iván Recasens Lahoz <irecasensl@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var raw bool

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Shows configured properties",
	Long:  `Exposes the configuration properties in a read-only way`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if raw {
			fmt.Print(GetConfigFileRaw())
		} else {
			fmt.Print(GetConfigFilePretty())
		}
	},
}

func init() {
	configCmd.AddCommand(showCmd)

	showCmd.Flags().BoolVarP(&raw, "raw", "r", false, "shows the raw content of the config file")
}
