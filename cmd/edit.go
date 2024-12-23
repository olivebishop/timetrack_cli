package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
     "timetrack/models"
    "strconv"
)

func newEditCmd() *cobra.Command {
    return &cobra.Command{
        Use:   "edit [task ID] [new name]",
        Short: "✏️ Edit a task's name",
        Args:  cobra.ExactArgs(2),
        RunE: func(cmd *cobra.Command, args []string) error {
            id, err := strconv.Atoi(args[0])
            if err != nil {
                return fmt.Errorf("❌ invalid task ID: %w", err)
            }

            if err := models.GetTaskManager().EditTask(id, args[1]); err != nil {
                return fmt.Errorf("❌ failed to edit task: %w", err)
            }
            fmt.Printf("✅ Task %d updated to '%s'!\n", id, args[1])
            return nil
        },
    }
}