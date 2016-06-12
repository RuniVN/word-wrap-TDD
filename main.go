package main

import "strings"

// WordWrap seperate paragraph with given length by
// adding \n to end of line
func WordWrap(para string, length int) string {
	listWords := strings.Split(para, " ")
	return strings.Join(listWords, "\n")
}
