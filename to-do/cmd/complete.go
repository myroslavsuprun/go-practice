package cmd

import (
	"strconv"
	"to-do/service"

	"github.com/spf13/cobra"
)

func completeCmd(action service.Complete) *cobra.Command {
	return &cobra.Command{

		Use:   "complete",
		Short: "Compete a to-do task",
		Long:  `Complete a to-do task from the list with a unique ID.`,
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				cmd.Print("Please provide a valid ID.")
			}

			title, err := action(id)
			if err != nil {
				cmd.Print("Failed to complete the task.")
				return
			}

			cmd.Print("Task completed successfully:", title)
		},
	}
}
