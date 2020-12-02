package sieve

import (
	"testing"
)

func TestEvenNumber(t *testing.T) {
	want := false
	got, _ := Sieve(42)

	if got != want {
		t.Errorf("Got %t want %t", got, want)
	}
}

func TestNegativeNumber(t *testing.T) {
	want := ErrNegativeNumber
	_, err := Sieve(-1)

	if err != want {
		t.Errorf("Expected error %s got %s", err, want)
	}
}
