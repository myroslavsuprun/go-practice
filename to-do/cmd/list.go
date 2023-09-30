package cmd

import (
	"to-do/service"

	"github.com/spf13/cobra"
)

func listCmd(action service.Get) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all the to-do tasks",
		Long:  `List all the to-do tasks with their unique IDs.`,
		Run: func(cmd *cobra.Command, args []string) {
			todos, err := action()

			if err != nil {
				cmd.Print("Failed to get todos.")
			}

			if len(todos) == 0 {
				cmd.Print("Your to-do list is empty. Lucky you!")
			}

			for _, todo := range todos {
				cmd.Printf("- %d: %s\n", todo.Id, todo.Title)
			}
		},
	}
}
