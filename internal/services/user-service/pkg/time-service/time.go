package timeService

import (
	"strconv"
	"strings"
	"time"
)

func ParseDuration(input string) (time.Duration, error) {
	if strings.HasSuffix(input, "d") {
		dayString := strings.TrimSuffix(input, "d")
		days, err := strconv.Atoi(dayString)
		if err != nil {
			return 0, ErrInvalidNumberOfDays
		}
		return time.Duration(days) * 24 * time.Hour, nil
	}

	return time.ParseDuration(input)
}
