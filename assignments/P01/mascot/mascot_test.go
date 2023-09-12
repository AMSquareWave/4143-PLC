package mascot_test

import (
	"testing"

	"example.com/go-demo-1/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
		t.Fatal("That ain't right.")
	} else if mascot.SecondBestMascot() != "Tux" {
		t.Fatal("You forgot about Tux.")
	}
}