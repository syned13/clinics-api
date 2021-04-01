package utils

import "regexp"

var hourRegex = regexp.MustCompile(`^([0-9]|0[0-9]|1[0-9]|2[0-3]):([0-9]|[0-5][0-9])$`)

// ValidateHour validates a valid hour in the format of hh:mm
func ValidateHour(hour string) bool {
	return hourRegex.MatchString(hour)
}
