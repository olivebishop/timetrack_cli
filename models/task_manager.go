package models

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type TaskManager struct {
	Tasks []Task `json:"tasks"`
	mu    sync.RWMutex
}

var (
	instance *TaskManager
	once     sync.Once
)

func GetTaskManager() *TaskManager {
	once.Do(func() {
		instance = &TaskManager{}
		instance.loadTasks()
	})
	return instance
}

func (tm *TaskManager) loadTasks() error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	dir := filepath.Join(configDir, "timetrack")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	file := filepath.Join(dir, "tasks.json")
	data, err := os.ReadFile(file)
	if err != nil {
		if os.IsNotExist(err) {
			tm.Tasks = []Task{}
			return nil
		}
		return err
	}

	return json.Unmarshal(data, &tm)
}

// Fixed saveTasks to not acquire additional locks
func (tm *TaskManager) saveTasks() error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	dir := filepath.Join(configDir, "timetrack")
	file := filepath.Join(dir, "tasks.json")

	data, err := json.MarshalIndent(tm, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(file, data, 0644)
}

func (tm *TaskManager) AddTask(task *Task) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	task.ID = len(tm.Tasks) + 1
	tm.Tasks = append(tm.Tasks, *task)
	return tm.saveTasks()
}

func (tm *TaskManager) StopTask(id int) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	for i := range tm.Tasks {
		if tm.Tasks[i].ID == id && !tm.Tasks[i].IsCompleted {
			tm.Tasks[i].EndTime = time.Now()
			tm.Tasks[i].Duration = tm.Tasks[i].EndTime.Sub(tm.Tasks[i].StartTime).Hours()
			tm.Tasks[i].IsCompleted = true
			return tm.saveTasks()
		}
	}
	return fmt.Errorf("task not found or already completed")
}

func (tm *TaskManager) EditTask(id int, newName string) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	for i := range tm.Tasks {
		if tm.Tasks[i].ID == id {
			tm.Tasks[i].Name = newName
			return tm.saveTasks()
		}
	}
	return fmt.Errorf("task not found")
}

func (tm *TaskManager) DeleteTask(id int) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	for i := range tm.Tasks {
		if tm.Tasks[i].ID == id {
			tm.Tasks = append(tm.Tasks[:i], tm.Tasks[i+1:]...)
			return tm.saveTasks()
		}
	}
	return fmt.Errorf("task not found")
}

func (tm *TaskManager) ListTasks() []Task {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	tasks := make([]Task, len(tm.Tasks))
	copy(tasks, tm.Tasks)
	return tasks
}
