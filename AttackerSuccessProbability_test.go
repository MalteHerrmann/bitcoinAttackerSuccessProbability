package main

import (
	"fmt"
	"testing"
)

func TestPoissonDensity(t *testing.T) {
	gotValue := poissonDensity(2.0, 0)
	got := fmt.Sprintf("%.3f", gotValue)
	want := "0.135"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestAttackerSuccessProbabilityShouldReturn1(t *testing.T) {
	got := AttackerSuccessProbability(0.3, 0)
	want := 1.0

	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestAttackerSuccessProbabilityShouldReturnFloat(t *testing.T) {
	gotValue := AttackerSuccessProbability(0.3, 5)
	got := fmt.Sprintf("%.7f", gotValue)
	want := "0.1773523"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
