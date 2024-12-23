package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "timetrack/models"
)

func newStartCmd() *cobra.Command {
    return &cobra.Command{
        Use:   "start [task name]",
        Short: "▶️ Start a new task",
        Args:  cobra.ExactArgs(1),
        RunE: func(cmd *cobra.Command, args []string) error {
            task := models.NewTask(args[0])
            if err := models.GetTaskManager().AddTask(task); err != nil {
                return fmt.Errorf("❌ failed to start task: %w", err)
            }
            fmt.Printf("✅ Task '%s' started!\n", args[0])
            return nil
        },
    }
}