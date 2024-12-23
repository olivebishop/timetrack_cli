package utils

import (
    "fmt"
    "github.com/fogleman/gg"
    "timetrack/models"
    "time"
)

type ReportGenerator struct {
    width  int
    height int
    dc     *gg.Context
}

func NewReportGenerator(width, height int) *ReportGenerator {
    return &ReportGenerator{
        width:  width,
        height: height,
        dc:     gg.NewContext(width, height),
    }
}

func (rg *ReportGenerator) GenerateReport(tasks []models.Task, outputPath string) error {
    // Set background
    rg.dc.SetRGB(0.95, 0.95, 0.95)
    rg.dc.Clear()

    // Try to load font, fallback to built-in if fails
    if err := rg.dc.LoadFontFace("/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf", 24); err != nil {
        if err := rg.dc.LoadFontFace("", 24); err != nil {
            return fmt.Errorf("failed to load any font: %w", err)
        }
    }

    // Draw header
    rg.dc.SetRGB(0.2, 0.2, 0.2)
    title := fmt.Sprintf("Daily Time Report - %s", time.Now().Format("2006-01-02"))
    rg.dc.DrawString(title, 50, 50)

    // Load regular font for task list
    if err := rg.dc.LoadFontFace("/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf", 16); err != nil {
        if err := rg.dc.LoadFontFace("", 16); err != nil {
            return fmt.Errorf("failed to load any font: %w", err)
        }
    }

    y := 100.0
    totalDuration := 0.0

    for _, task := range tasks {
        if task.IsCompleted {
            // Draw task box
            rg.dc.SetRGB(0.9, 0.9, 0.9)
            rg.dc.DrawRoundedRectangle(40, y-20, float64(rg.width)-80, 40, 10)
            rg.dc.Fill()

            // Draw task text
            rg.dc.SetRGB(0.2, 0.2, 0.2)
            taskText := fmt.Sprintf("%s - %.2f hours", task.Name, task.Duration)
            rg.dc.DrawString(taskText, 50, y)
            
            y += 50
            totalDuration += task.Duration
        }
    }

    // Draw total time
    rg.dc.SetRGB(0.1, 0.1, 0.1)
    if err := rg.dc.LoadFontFace("/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf", 20); err != nil {
        if err := rg.dc.LoadFontFace("", 20); err != nil {
            return fmt.Errorf("failed to load any font: %w", err)
        }
    }
        rg.dc.DrawString(fmt.Sprintf("Total Time: %.2f hours", totalDuration), 50, y)
        return rg.dc.SavePNG(outputPath)
    }