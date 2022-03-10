package utils

import (
	"fmt"
	"os"
	"time"
)

func HandleError(err interface{}) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "An error occured:", err)
		os.Exit(1)
	}
}

func UnixToDate(unix int64) string {
	t := time.Unix(unix, 0)
	layout := "2006-01-02 15:04:05"
	return t.Format(layout)
}
