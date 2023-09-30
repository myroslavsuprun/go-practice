package cmd

import (
	"to-do/service"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "This is a to-do CLI app",
	Long:  `This is a to-do CLI app in which you can add, remove, complete and list tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute(actions service.Actions) error {
	rootCmd.AddCommand(addCmd(actions.Add), listCmd(actions.Get))
	rootCmd.AddCommand(completeCmd(actions.Completed), clearCmd(actions.Clear))
	rootCmd.AddCommand(rmCmd(actions.Remove), todayCmd(actions.GetToday))
	return rootCmd.Execute()
}
