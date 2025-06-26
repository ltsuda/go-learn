package iteration

import "strings"

const repeatCount = 5

func Repeat(character string, maxCount int) string {
	var repeated strings.Builder
	var maxRepeat = repeatCount

	if maxCount > 0 {
		maxRepeat = maxCount
	}

	for i := 0; i < maxRepeat; i++ {
		repeated.WriteString(character)
	}

	return repeated.String()
}
