package main

import "testing"

func TestPrompt(t *testing.T) {
	t.Run("asking about entering principal", func(t *testing.T) {
		got := Prompt("Enter the principal: ")
		want := "Enter the principal: " 

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}