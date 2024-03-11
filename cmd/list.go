/*
Copyright © 2024 Iván Recasens Lahoz
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var printAll bool
var printOnlyCompleted bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists tasks",
	Long: `List the tasks.
By default it only shows pending tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.Open(GetVaultDirectory() + "\\" + GetFilename())
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		// Create a new scanner
		scanner := bufio.NewScanner(file)

		// Read lines
		for scanner.Scan() {
			current := scanner.Text()
			finished := strings.Contains(current, "- [x]")
			// fmt.Printf("%s is finished? %v\n", current, finished)

			if hasToBePrinted(finished) {
				fmt.Println(current)
			}
		}
	},
}

func hasToBePrinted(finished bool) bool {
	if printAll {
		return true
	}

	if printOnlyCompleted {
		return finished
	}

	if !finished {
		return true
	}

	return false
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&printAll, "all", "a", false, "shows both pending and completed TODOs")
	listCmd.Flags().BoolVarP(&printOnlyCompleted, "finished", "f", false, "shows only completed TODOs")
}
