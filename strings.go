package main

import (
	"strings"
	"time"
)

// Capitalize a string
func Capitalize(value string) string {
	// Treat dashes as spaces
	value = strings.Replace(value, "-", " ", 99)
	valuesplit := strings.Split(value, " ")
	var result string
	for _, v := range valuesplit {
		if(len(v) <= 0) {
			continue
		}
		result += strings.ToUpper(v[:1])
		result += v[1:] + " "
	}
	return result
}

// Trim a string to 128 characters, for meta tags.
func TrimForMeta(value string) string {
	if(len(value) <= 127) {
		return value
	}
	return value[:128]+"..."
}

// Print the server date three months from now 
func PrintThreeMonthsFromNow() string {
	future := time.Now().Add(time.Hour*2190)
	return future.Format("Jan 02 2006, 03:04:05PM -0700")
}