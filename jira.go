package main

import (
	"github.com/andygrunwald/go-jira"
	"os"
)

type Client jira.Client

func getJiraClient(login, token string) (*Client, error) {
	tp := jira.BasicAuthTransport{
		Username: login,
		Password: token,
	}

	client, error := jira.NewClient(tp.Client(), os.Getenv("JIRA_URL"))

	return (*Client)(client), error
}

func (c *Client) addWorkLog(issueId, description string, timeSpent int64, startTime jira.Time) (*jira.WorklogRecord, error) {
	timeSpentSeconds := int(timeSpent)

	worklogRecord := jira.WorklogRecord{
		Comment:          description,
		Started:          &startTime,
		TimeSpentSeconds: timeSpentSeconds,
	}

	worklog, _, err := c.Issue.AddWorklogRecord(issueId, &worklogRecord)

	return worklog, err
}

func (c *Client) updateWorkLog(issueId, description, worklogId string, timeSpent int64, startTime jira.Time) (*jira.WorklogRecord, error) {
	timeSpentSeconds := int(timeSpent)

	worklogRecord := jira.WorklogRecord{
		ID:               worklogId,
		Comment:          description,
		Started:          &startTime,
		TimeSpentSeconds: timeSpentSeconds,
	}

	worklog, _, err := c.Issue.AddWorklogRecord(issueId, &worklogRecord)

	return worklog, err
}
