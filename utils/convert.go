package util

import (
	"log"
	"strconv"
)

// StringToInt converts a string to int
func StringToInt(s string) int {
	v, err := strconv.Atoi(s)
	if err == nil {
		return v
	}
	log.Fatal("Unable to convert: ", s, " to int: ", err)
	return -1
}
