# go-toggl-jira-sync
Sync toggl time entries with jira task worklogs

## How to use
1. Put url to Jira in `.env` or `.env.local` file
2. Create file `users.json` from `users.json.dist` (`$cp users.json.dist users.json`) and fill `JIRA_URL` variable
3. Run  `go run .` or build `go build .` project
