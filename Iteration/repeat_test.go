package iteration

import (
	"fmt"
	"strings"
	"testing"
)

const characterRepeatedTest = "a"
const repeatCountsTest = 5

func TestRepeat(t *testing.T) {
	repeated := Repeat(characterRepeatedTest)
	expected := strings.Repeat(characterRepeatedTest, repeatCountsTest)

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat(characterRepeatedTest)
	fmt.Println(repeated)
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(characterRepeatedTest)
	}
}
