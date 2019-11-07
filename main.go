package main

import (
	"flag"
	"fmt"
	"github.com/google/logger"
	"github.com/joho/godotenv"
	"github.com/recoilme/slowpoke"
	"regexp"
	"strconv"
	"strings"

	//"github.com/andygrunwald/go-jira"
	"os"
)

const logPath = "./app.log"

func init() {
	// Load and overwrite variables from local envirement file
	godotenv.Load(".env.local", ".env")
}

func main() {
	flag.Parse()

	lf, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}

	defer slowpoke.CloseAll()
	defer lf.Close()

	defer logger.Init("Logger", false, true, lf).Close()

	for _, user := range getUsers() {
		account, _ := getTogglAccount(user.TogglToken)

		for _, entry := range account.getTogglEntries() {
			if duration := getEntryFromDB(entry.Id); strconv.Itoa(duration) != strconv.FormatInt(entry.Duration, 10) {
				setEntryInDB(entry.Id, entry.Duration)

				//TODO
			}

			//TODO move above in if statement
			client, _ := getJiraClient(user.JiraLogin, user.JiraToken)

			//TODO move to jira
			issueId := strings.Split(entry.Task.Name, " ")
			issues, _, _ := client.Issue.Search("text ~ '" + issueId[0] + "' order by lastViewed DESC", nil)
			fmt.Println(issues)
			re := regexp.MustCompile(`\w+\-\d*`)
			re.FindAllSubmatch([]byte(entry.Task.Name), 1)

			//	now := &jira.Time{}
			//
			//	//worklogRecord := jira.WorklogRecord{"", u, nil, "test", now, now, now, "1h", 0, "", issue.ID, nil}
			//	_, _, err := client.Issue.AddWorklogRecord(issue.ID, &worklogRecord)
		}
	}
}
