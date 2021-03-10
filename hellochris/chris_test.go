package main

import "testing"

func TestChris(t *testing.T) {

	assertCorrectMsg := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("in spanish", func(t *testing.T) {
		got := Chris("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMsg(t, got, want)
	})
	t.Run("saying hello to people", func(t *testing.T) {
		got := Chris("Chris", "English")
		want := "Hello, Chris"

		assertCorrectMsg(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Chris("", "")
		want := "Hello, World"

		assertCorrectMsg(t, got, want)
	})

}
