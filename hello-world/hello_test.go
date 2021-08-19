package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		assertCorrectMessage(t, Hello("Chris", ""), "Hello Chris!!")
	})
	t.Run(`say "Hello World!" when empty string is supplied`, func(t *testing.T) {
		assertCorrectMessage(t, Hello("", ""), "Hello World!!")
	})
	t.Run("in Spanish", func(t *testing.T) {
		assertCorrectMessage(t, Hello("Elodie", "Spanish"), "Hola Elodie!!")
	})
	t.Run("in French", func(t *testing.T) {
		assertCorrectMessage(t, Hello("Renee", "French"), "Bonjour Renee!!")
	})
}
