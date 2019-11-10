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

func getTogglSession(token string) toggl.Session {
	return toggl.OpenSession(token)
}

// Return entries with Project and Task only
func (s *Session) getTogglEntries(days int) []TogglEntry {
	if days == 0 {
		days = 7
	}

	var entries []TogglEntry

	e, _ := s.GetTimeEntries(time.Now().AddDate(0, 0, -days), time.Now())

	for _, entry := range e {
		project, err := s.getTogglProjectById(entry.Pid)

		if err != nil {
			logger.Infof("[toggl] Entry '%s' with id %d dosen't have project", entry.Description, entry.ID)
			continue
		}

		task, err := s.getTogglTaskById(entry.Tid)

		if err != nil {
			logger.Infof("[toggl] Entry '%s' with id %d dosen't have task", entry.Description, entry.ID)
			continue
		}

		entry := TogglEntry{
			Id:          entry.ID,
			Start:       *entry.Start,
			Stop:        *entry.Stop,
			Duration:    entry.Duration,
			Description: entry.Description,
			Project:     project,
			Task:        task,
		}

		entries = append(entries, entry)
	}

	return entries
}

func (s *Session) getTogglProjectById(id int) (TogglProject, error) {
	for _, project := range s.getTogglProjects() {
		if project.Id == id {
			return project, nil
		}
	}

	return TogglProject{}, fmt.Errorf("Project with id %d not exists", id)
}

func (s *Session) getTogglProjects() []TogglProject {
	var projects []TogglProject

	account, _ := s.GetAccount()

	for _, project := range account.Data.Projects {

		project := TogglProject{
			Id:   project.ID,
			Name: project.Name,
		}

		projects = append(projects, project)
	}

	return projects
}

func (s *Session) getTogglTaskById(id int) (TogglTask, error) {
	for _, task := range s.getTogglTasks() {
		if task.Id == id {
			return task, nil
		}
	}

	return TogglTask{}, fmt.Errorf("Task with id %d not exists", id)
}

func (s *Session) getTogglTasks() []TogglTask {
	var tasks []TogglTask
	account, _ := s.GetAccount()

	for _, task := range account.Data.Tasks {
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
