package main

import (
	"fmt"
	"github.com/google/logger"
	"github.com/recoilme/slowpoke"
	"os"
	"strconv"
)

const dbFile = "db/database.db"

func setEntryInDB(id int, duration int64) error {
	err := slowpoke.Set(dbFile, []byte(strconv.Itoa(id)), []byte(strconv.Itoa(int(duration))))

	if err != nil {
		logger.Fatal(err)
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}

	return err
}

func getEntryFromDB(id int) int {
	res, err := slowpoke.Get(dbFile, []byte(strconv.Itoa(id)))

	if err != nil {
		return 0
	}

	duration, err := strconv.Atoi(string(res))
	return duration
}