package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		// db
		fmt.Printf("Added \"%s\" to your task list.\n")
	},
}
