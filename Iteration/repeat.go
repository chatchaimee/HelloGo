package iteration

const repeatCounts = 5

func Repeat(character string) string {
	repeated := ""
	for i := 0; i < repeatCounts; i++ {
		repeated += character
	}
	return repeated
}
