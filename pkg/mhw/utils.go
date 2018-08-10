package mhw

import (
	"log"
	"strconv"
)

func checkError(e error) {
	if e != nil {
		log.Fatalf("error: %v", e)
		panic(e)
	}
}

func parseCheckedInt(text string) int {
	value, err := strconv.Atoi(text)
	if err != nil {
		log.Fatalf("error: %v", err)
		panic(err)
	}
	return value
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
