package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	/* 	var drow []string
	   	for _, r := range word {
	   		drow = append([]string{string(r)}, drow...)
	   	}
	   	return strings.Join(drow, "") */

	var drow string
	for _, r := range word {
		drow = string(r) + drow
	}
	return drow
}
