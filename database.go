package main

import (
	"fmt"
	"github.com/google/logger"
	"github.com/recoilme/slowpoke"
	"os"
	"strconv"
)

const dbFile = "db/database.db"

func setEntryInDB(id int, value string) error {
	err := slowpoke.Set(dbFile, []byte(strconv.Itoa(id)), []byte(value))

	if err != nil {
		logger.Fatal(err)
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}

	return err
}

func getEntryFromDB(id int) string {
	res, err := slowpoke.Get(dbFile, []byte(strconv.Itoa(id)))

	if err != nil {
		return "0 0"
	}

	return string(res)
}
