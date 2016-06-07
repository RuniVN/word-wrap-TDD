package main

// WordWrap seperate paragraph with given length by
// adding \n to end of line
func WordWrap(para string, length int) string {
	var result string
	for i, v := range para {
		if (i+1)%length == 0 {
			result += string(v) + "\n"
		} else {
			result += string(v)
		}
	}

	return result
}
