package cmd

import (
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "timetrack",
    Short: "⏰ A Time tracking CLI application",
    Long: `⏰ TimeTrack CLI - A simple and efficient time tracking tool
Track your daily tasks, monitor time spent, and generate beautiful reports.`,
}

func Execute() error {
    return rootCmd.Execute()
}

func init() {
    rootCmd.AddCommand(
        newStartCmd(),
        newStopCmd(),
        newListCmd(),
        newEditCmd(),
        newDeleteCmd(),
        newReportCmd(),
    )
}
