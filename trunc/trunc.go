package trunc

import (
	"strconv"
)

// Trunc func take string input and return only number part.
func Trunc(input string) (int, error) {
	// Change string input to floating number
	floatValue, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}

	// Truncate to integer
	intValue := int(floatValue)
	return intValue, nil
}
