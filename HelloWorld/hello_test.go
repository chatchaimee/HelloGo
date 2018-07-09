package main

import "testing"

const name = "Jay"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("saying hello to people in English", func(t *testing.T) {
		got := Hello(name, "")
		want := "Hello!, " + name

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello!, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in Spanish", func(t *testing.T) {
		got := Hello(name, "Spanish")
		want := "Hola, " + name
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in French", func(t *testing.T) {
		got := Hello(name, "French")
		want := "Bonjour, " + name
		assertCorrectMessage(t, got, want)
	})
}
