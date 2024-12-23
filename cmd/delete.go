package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "timetrack/models"
    "strconv"
)

func newDeleteCmd() *cobra.Command {
    return &cobra.Command{
        Use:   "delete [task ID]",
        Short: "🗑️ Delete a task",
        Args:  cobra.ExactArgs(1),
        RunE: func(cmd *cobra.Command, args []string) error {
            id, err := strconv.Atoi(args[0])
            if err != nil {
                return fmt.Errorf("❌ invalid task ID: %w", err)
            }

            if err := models.GetTaskManager().DeleteTask(id); err != nil {
                return fmt.Errorf("❌ failed to delete task: %w", err)
            }
            fmt.Printf("✅ Task %d deleted!\n", id)
            return nil
        },
    }
}
