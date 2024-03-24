// This package provide a single function to generate the radio validation code for Renault cars.
package codeautoradio

import (
	"fmt"
)

// precodeIsValid checks if the given precode is valid.
func precodeIsValid(precode string) bool {
	return len(precode) == 4 && // must be 4 characters long
		'A' <= precode[0] && precode[0] <= 'Z' && // must start with a letter
		(precode[0] != 'A' || precode[1] != '0') && // must not start with A0
		precode[1] >= '0' && precode[1] <= '9' && // the second character must be a digit
		precode[2] >= '0' && precode[2] <= '9' && // the third character must be a digit
		precode[3] >= '0' && precode[3] <= '9' // the fourth character must be a digit
}

// GetRadioCode returns the radio code for the given precode.
// The precode must be a string of 4 characters.
// The first character must be an upper case letter and the following 3 characters must be digits.
// The precode must not start with A0.
// If the precode is not valid, the function returns an empty string and false.
// Otherwise, it returns the radio (validation) code as a string and true.
func GetRadioCode(precode string) (code string, ok bool) {
	if !precodeIsValid(precode) {
		return "", false
	}
	x := int(precode[1]) + int(precode[0])*10 - 698
	y := int(precode[3]) + int(precode[2])*10 + x - 528
	z := (y * 7) % 100
	c := z/10 + z%10*10 + ((259%x)%100)*100
	return fmt.Sprintf("%04d", c), true
}
