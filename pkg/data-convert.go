package pkg

import (
	"strconv"
	"strings"
)

func FromStringToPositiveNumber(str string) float64 {
	cleanStr := strings.ReplaceAll(str, "$", "")
	cleanStr = strings.TrimSpace(cleanStr)

	number, err := strconv.ParseFloat(cleanStr, 64)

	if err != nil || number <= 0 {
		number = 0
	}

	return number
}
