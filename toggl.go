package main

import (
	"fmt"
	"github.com/google/logger"
	"github.com/jason0x43/go-toggl"
	"regexp"
	"time"
)

type TogglProject struct {
	Id   int
	Name string
}

type TogglTask struct {
	Id     int
	JiraId string
	Name   string
}

type TogglEntry struct {
	Id          int
	Start       time.Time
	Stop        time.Time
	Duration    int64
	Description string
	Project     TogglProject
	Task        TogglTask
	Imported    bool
}

type Session struct {
	toggl.Session
}

type Account struct {
	toggl.Account
}

func getTogglSession(token string) toggl.Session {
	return toggl.OpenSession(token)
}

// Return entries with Project and Task only
func (s *Session) getTogglEntries(days int) []TogglEntry {
	if days == 0 {
		days = 7
	}

	a, _ := s.GetAccount()
	account := Account{a}

	var entries []TogglEntry

	e, _ := s.GetTimeEntries(time.Now().AddDate(0, 0, -days), time.Now())

	for _, entry := range e {
		project, error := account.getTogglProjectById(entry.Pid)

		if error != nil {
			logger.Infof("[toggl] Entry '%s' with id %d dosen't have project", entry.Description, entry.ID)
			continue
		}

		task, error := account.getTogglTaskById(entry.Tid)

		if error != nil {
			logger.Infof("[toggl] Entry '%s' with id %d dosen't have task", entry.Description, entry.ID)
			continue
		}

		entry := TogglEntry{
			Id:          entry.ID,
			Start:       entry.Start.Add(time.Hour).Add(time.Millisecond),
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
	for _, task := range a.getTogglTasks() {
		if task.Id == id {
			return task, nil
		}
	}

	return TogglTask{}, fmt.Errorf("Task with id %d not exists", id)
}

func (a *Account) getTogglTasks() []TogglTask {
	var tasks []TogglTask

	for _, task := range a.Data.Tasks {
		jiraId := ""

		regex := regexp.MustCompile(`\w+\-\d*`)
		regexResults := regex.FindAllSubmatch([]byte(task.Name), 1)

		if len(task.Name) > 0 && len(regexResults) > 0 {
			jiraId = string(regexResults[0][0])
		}

		task := TogglTask{
			Id:     task.ID,
			JiraId: jiraId,
			Name:   task.Name,
		}

		tasks = append(tasks, task)
	}

	return tasks
}
