package cmd

import (
	"strings"
	"to-do/service"

	"github.com/spf13/cobra"
)

func addCmd(action service.Add) *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Add a new to-do task",
		Long:  `Add a new to-do task to the list with a unique ID.`,
		Run: func(cmd *cobra.Command, args []string) {
			taskTitle := strings.Join(args, " ")

			err := action(taskTitle)
			if err != nil {
				cmd.Print("Failed to add the task.")
				return
			}
			cmd.Print("Task added successfully:", taskTitle)
		},
	}
}
