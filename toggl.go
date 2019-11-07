package main

import (
	"fmt"
	g "github.com/jason0x43/go-toggl"
	"github.com/google/logger"
	"time"
)

type TogglProject struct {
	Id   int
	Name string
}

type TogglTask struct {
	Id   int
	Name string
}

type TogglEntry struct {
	Id          int
	Start		time.Time
	Stop        time.Time
	Duration    int64
	Description string
	Project     TogglProject
	Task        TogglTask
	Imported    bool
}

type Account g.Account

func getTogglAccount(token string) (Account, error) {
	togglSession := g.OpenSession(token)
	account, error := togglSession.GetAccount()

	return Account(account), error
}

// Return entries with Project and Task only
func (a *Account) getTogglEntries() []TogglEntry {
	var entries []TogglEntry

	for _, entry := range a.Data.TimeEntries {
		project, err := a.getTogglProjectById(entry.Pid)

		if err != nil {
			logger.Infof("[toggl] Entry '%s' with id %d dosen't have project", entry.Description, entry.ID)
			continue
		}

		task, err := a.getTogglTaskById(entry.Tid)

		if err != nil {
			logger.Infof("[toggl] Entry '%s' with id %d dosen't have task", entry.Description, entry.ID)
			continue
		}

		entry := TogglEntry{
			Id:          entry.ID,
			Start:		 *entry.Start,
			Stop:		 *entry.Stop,
			Duration:    entry.Duration,
			Description: entry.Description,
			Project:     project,
			Task:        task,
		}

		entries = append(entries, entry)
	}

	return entries
}

func (a *Account) getTogglProjectById(id int) (TogglProject, error) {
	for _, project := range a.getTogglProjects() {
		if project.Id == id {
			return project, nil
		}
	}

	return TogglProject{}, fmt.Errorf("Project with id %d not exists", id)
}

func (a *Account) getTogglProjects() []TogglProject {
	var projects []TogglProject

	for _, project := range a.Data.Projects {
		project := TogglProject{
			Id:   project.ID,
			Name: project.Name,
		}

		projects = append(projects, project)
	}

	return projects
}

func (a *Account) getTogglTaskById(id int) (TogglTask, error) {
	for _, task:= range a.getTogglTasks() {
		if task.Id == id {
			return task, nil
		}
	}

	return TogglTask{}, fmt.Errorf("Task with id %d not exists", id)
}

func (a *Account) getTogglTasks() []TogglTask {
	var tasks []TogglTask

	for _, task := range a.Data.Tasks {
		task := TogglTask{
			Id:   task.ID,
			Name: task.Name,
		}

		tasks = append(tasks, task)
	}

	return tasks
}
