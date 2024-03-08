/*
Copyright © 2024 Iván Recasens Lahoz <irecasensl@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var fileDirectory string
var filename string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new task to the todo list",
	Long:  `Adds the specified task to the todo list`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		todo := args[0]
		path := fmt.Sprintf("%s\\%s", GetVaultDirectory(), filename)

		file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
		if err != nil {
			panic("Failed to read todo list")
		}
		defer file.Close()

		file.WriteString("- [ ] " + todo + "\n")

		fmt.Printf("Added \"%s\" to the todo list\n", todo)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&fileDirectory, "directory", "d", GetVaultDirectory(), "Specifies the directory of the file")
	addCmd.Flags().StringVarP(&filename, "file", "f", "TODO.md", "Specifies which file to use")
}
