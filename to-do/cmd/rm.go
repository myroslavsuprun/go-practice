package cmd

import (
	"strconv"
	"to-do/service"

	"github.com/spf13/cobra"
)

func rmCmd(action service.Remove) *cobra.Command {
	return &cobra.Command{
		Use:   "rm",
		Short: "Remove a to-do task",
		Long:  `Remove a to-do task from the list with a unique ID.`,
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				cmd.Print("Please provide a valid ID.")
				return
			}

			title, err := action(id)
			if err != nil {
				cmd.Print("The to-do not found.")
				return
			}

			cmd.Print("The to-do has been removed:", title)
		},
	}
}
