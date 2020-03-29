package module01

import (
	"fmt"
	"log"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	var res int
	mult := 1
	for i := len(value) - 1; i >= 0; i-- {
		var val int
		_, err := fmt.Sscanf(string(value[i]), "%X", &val)
		if err != nil {
			log.Panicf("Unable to translate value: %v", value[i])
		}
		res += mult * val
		mult *= base
	}
	return res
}
