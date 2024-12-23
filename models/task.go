package models

import (
    "time"
)

type Task struct {
    ID          int       `json:"id"`
    Name        string    `json:"name"`
    StartTime   time.Time `json:"start_time"`
    EndTime     time.Time `json:"end_time,omitempty"`
    Duration    float64   `json:"duration,omitempty"`
    IsCompleted bool      `json:"is_completed"`
}

func NewTask(name string) *Task {
    return &Task{
        Name:      name,
        StartTime: time.Now(),
    }
}