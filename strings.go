package main

import (
	"strings"
)

// Capitalize a string
func Capitalize(value string) string {
	// Treat dashes as spaces
	value = strings.Replace(value, "-", " ", 99)
	valuesplit := strings.Split(value, " ")
	var result string
	for _, v := range valuesplit {
		result += strings.ToUpper(v[:1])
		result += v[1:] + " "
	}
	return result
}