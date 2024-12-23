package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "timetrack/models"
    "timetrack/utils"
    "path/filepath"
    "time"
)

func newReportCmd() *cobra.Command {
    return &cobra.Command{
        Use:   "report",
        Short: "📊 Generate a daily report",
        RunE: func(cmd *cobra.Command, args []string) error {
            tasks := models.GetTaskManager().ListTasks()
            
            filename := fmt.Sprintf("timetrack-report-%s.png", 
                time.Now().Format("2006-01-02"))
            
            reportGen := utils.NewReportGenerator(800, 600)
            if err := reportGen.GenerateReport(tasks, filename); err != nil {
                return fmt.Errorf("❌ failed to generate report: %w", err)
            }
            
            absPath, _ := filepath.Abs(filename)
            fmt.Printf("✅ Report generated: %s\n", absPath)
            return nil
        },
    }
}