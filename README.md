# go-toggl-jira-sync
Sync toggl time entries with jira task worklogs

## How to use
1. Create .env file from .env.dist (`$cp users.json.dist users.json`) and put url to Jira
2. Create file `users.json` from `users.json.dist` (`$cp users.json.dist users.json`) and fill `JIRA_URL` variable
3. Run  `go run .` or build `go build .` project

## TODO
- Allow to select project to sync
- Statistics after sync process
- GUI
