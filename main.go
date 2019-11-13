package main

import (
	"flag"
	"github.com/andygrunwald/go-jira"
	"github.com/google/logger"
	"github.com/joho/godotenv"
	"github.com/recoilme/slowpoke"
	"os"
	"strconv"
	"strings"
)

const logPath = "./app.log"

func init() {
	// Load and overwrite variables from local envirement file
	godotenv.Load(".env.local", ".env")
}

func main() {
	flag.Parse()

	lf, error := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if error != nil {
		logger.Fatalf("Failed to open log file: %v", error)
	}

	defer slowpoke.CloseAll()
	defer lf.Close()

	defer logger.Init("Logger", false, true, lf).Close()

	for _, user := range getUsers() {
		togglSession := getTogglSession(user.TogglToken)
		session := &Session{togglSession}
		jiraClient, _ := getJiraClient(user.JiraLogin, user.JiraToken)

		days, _ := strconv.Atoi(os.Getenv("DAYS"))

		for _, entry := range session.getTogglEntries(days) {
			value := getEntryFromDB(entry.Id)
			duration := strings.Split(value, " ")[0]
			jiraWorklogId := strings.Split(value, " ")[1]

			description := ""

			if entry.Task.Name != entry.Description {
				description = entry.Description
			}

			if duration != strconv.FormatInt(entry.Duration, 10) && entry.Task.JiraId != "" {
				if jiraWorklogId != "0" {
					//update
					//TODO https://github.com/prugala/go-toggl-jira-sync/issues/2
					//start := jira.Time(entry.Start)
					//worklog, error := jiraClient.updateWorkLog(entry.Task.JiraId, description, jiraWorklogId, entry.Duration, start)
					//
					//if error == nil {
					//	setEntryInDB(entry.Id, strconv.FormatInt(entry.Duration, 10)+" "+worklog.ID)
					//} else {
					//	logger.Fatalf("[jira] Task %s error: %v", entry.Task.JiraId, error)
					//}
				} else {
					//new
					start := jira.Time(entry.Start)
					worklog, error := jiraClient.addWorkLog(entry.Task.JiraId, description, entry.Duration, start)

					if error == nil {
						setEntryInDB(entry.Id, strconv.FormatInt(entry.Duration, 10)+" "+worklog.ID)
					} else {
						logger.Errorf("[jira] Task %s error: %v", entry.Task.JiraId, error)
					}
				}
			}
		}
	}
}
