package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "timetrack/models"
    "strconv"
)

func newStopCmd() *cobra.Command {
    return &cobra.Command{
        Use:   "stop [task ID]",
        Short: "⏹️ Stop a running task",
        Args:  cobra.ExactArgs(1),
        RunE: func(cmd *cobra.Command, args []string) error {
            id, err := strconv.Atoi(args[0])
            if err != nil {
                return fmt.Errorf("❌ invalid task ID: %w", err)
            }

            if err := models.GetTaskManager().StopTask(id); err != nil {
                return fmt.Errorf("❌ failed to stop task: %w", err)
            }
            fmt.Printf("✅ Task %d stopped!\n", id)
            return nil
        },
    }
}