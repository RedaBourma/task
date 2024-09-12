package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"task/db"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreatTask(task)
		if err != nil {
			return
		}
		fmt.Printf("Added \"%s\" to your task list.\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
