package purchases

import (
	"regexp"
	"strconv"
)

const DefaultEstimatedDays uint = 7

var etdDigits = regexp.MustCompile(`\d+`)

func ParseETDDays(etd string) (days uint, ok bool) {
	var maxDays uint

	for _, match := range etdDigits.FindAllString(etd, -1) {
		n, err := strconv.ParseUint(match, 10, 32)
		if err != nil {
			continue
		}
		if uint(n) > maxDays {
			maxDays = uint(n)
		}
	}

	return maxDays, maxDays > 0
}
