package cmd

import (
	"to-do/service"

	"github.com/spf13/cobra"
)

func clearCmd(action service.Clear) *cobra.Command {
	return &cobra.Command{
		Use:   "clear",
		Short: "Clear all the to-do tasks",
		Long:  `Remove all the to-do tasks from the list and reset the ID count.`,
		Run: func(cmd *cobra.Command, args []string) {

			err := action()
			if err != nil {
				cmd.Print("Failed to clear the list.")
				return
			}

			cmd.Print("List has been cleared.")
		},
	}
}
