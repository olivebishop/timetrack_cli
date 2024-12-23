package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "timetrack/models"
)

func newListCmd() *cobra.Command {
    return &cobra.Command{
        Use:   "list",
        Short: "📋 List all tasks",
        RunE: func(cmd *cobra.Command, args []string) error {
            tasks := models.GetTaskManager().ListTasks()
            if len(tasks) == 0 {
                fmt.Println("📝 No tasks found")
                return nil
            }

            fmt.Println("📊 Tasks:")
            for _, task := range tasks {
                status := "🏃 In Progress"
                if task.IsCompleted {
                    status = fmt.Sprintf("✅ Completed (%.2f hours)", task.Duration)
                }
                fmt.Printf("ID: %d, Name: %s, Status: %s\n", task.ID, task.Name, status)
            }
            return nil
        },
    }
}
