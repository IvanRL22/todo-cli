/*
Copyright © 2024 Iván Recasens Lahoz <irecasensl@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new task to the todo list",
	Long:  `Adds the specified task to the todo list`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		todo := args[0]

		home, err := os.UserHomeDir()
		if err != nil {
			panic("Failed to obtain home directory")
		}

		file, err := os.OpenFile(home+"\\todoList.txt", os.O_CREATE, 0777)
		if err != nil {
			panic("Failed to read todo list")
		}
		defer file.Close()

		file.WriteString(todo + "\n")

		fmt.Printf("Added \"%s\" to the todo list\n", todo)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
