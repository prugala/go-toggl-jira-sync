package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/logger"
	"io/ioutil"
	"os"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	JiraName string `json:"jira_name"`
	JiraLogin string `json:"jira_login"`
	JiraToken string `json:"jira_token"`
	TogglToken string `json:"toggl_token"`
}

func getUsers() []User {
	var users Users

	jsonFile, err := os.Open("users.json")

	if err != nil {
		logger.Fatal(err)
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &users)

	defer jsonFile.Close()

	return users.Users
}