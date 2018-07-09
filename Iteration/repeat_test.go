package iteration

import (
	"fmt"
	"strings"
	"testing"
)

const characterRepeatedTest = "a"
const repeatCountsTest = 5

func TestRepeat(t *testing.T) {
	assertCorrect := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("Test Repeat Function", func(t *testing.T) {
		repeated := Repeat(characterRepeatedTest)
		expected := "aaaaa"

		assertCorrect(t, repeated, expected)
	})

	t.Run("Compare Repeat Function With Repeat Package Must be Equal", func(t *testing.T) {
		repeated := Repeat(characterRepeatedTest)
		expected := strings.Repeat(characterRepeatedTest, repeatCountsTest)

		assertCorrect(t, repeated, expected)
	})
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
