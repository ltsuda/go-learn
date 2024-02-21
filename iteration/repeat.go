package iteration

const repeatCount = 5

func Repeat(character string, maxCount int) string {
	var repeated string
	var maxRepeat int = repeatCount

	if maxCount > 0 {
		maxRepeat = maxCount
	}

	for i := 0; i < maxRepeat; i++ {
		repeated += character
	}

	return repeated
}
