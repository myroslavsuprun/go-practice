package cmd

import (
	"to-do/service"

	"github.com/spf13/cobra"
)

func todayCmd(action service.GetToday) *cobra.Command {
	return &cobra.Command{
		Use:   "today",
		Short: "Show completed today to-do tasks",
		Long:  `Show completed today to-do tasks.`,
		Run: func(cmd *cobra.Command, args []string) {
			todos, err := action()
			if err != nil {
				cmd.Print("Failed to get todos.")
				return
			}

			if len(todos) == 0 {
				cmd.Print("You haven't done anything today.")
			}

			for _, todo := range todos {
				cmd.Printf("- %d: %s\n", todo.Id, todo.Title)
			}
		},
	}
}
