package main

import "strings"

// WordWrap seperate paragraph with given length by
// adding \n to end of line
func WordWrap(para string, length int) string {
	if len(para) < length {
		return para
	}

	listWords := strings.Fields(para)
	var listWordsSanitized []string
	for _, v := range listWords {
		if len(v) <= length {
			listWordsSanitized = append(listWordsSanitized, v)
		} else {
			listWordsSanitized = append(listWordsSanitized, v[0:length], v[length:len(v)])
		}

	}
	return strings.Join(listWordsSanitized, "\n")
}
