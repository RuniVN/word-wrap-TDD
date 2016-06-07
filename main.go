package main

// WordWrap seperate paragraph with given length by
// adding \n to end of line
func WordWrap(para string, length int) string {
	var result string
	var j int
	for i := 0; i <= len(para)-1; i++ {
		if (j+1)%length == 0 {
			result += string(para[i]) + "\n"
		} else if string(para[i]) == " " {
			result += "\n"
			j = 0
			continue
		} else {
			result += string(para[i])
		}
		j++
	}

	return result
}
