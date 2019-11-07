package main

import (
	"github.com/andygrunwald/go-jira"
	"os"
)

func getJiraClient(login, token string) (*jira.Client, error) {
	tp := jira.BasicAuthTransport{
		Username: login,
		Password: token,
	}

	return jira.NewClient(tp.Client(), os.Getenv("JIRA_URL"))
}