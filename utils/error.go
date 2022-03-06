package utils

import (
	"fmt"
	"os"
)

func HandleError(err interface{}) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
