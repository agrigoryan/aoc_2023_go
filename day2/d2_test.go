package day2

import (
	"testing"

	"github.com/agrigoryan/aoc_2023_go/aocutils"
)

func TestD2p1(t *testing.T) {
	want := 2377
	if r := d2p1(aocutils.ReadInput("d2p1.txt")); r != want {
		t.Errorf("d2p1(..) got %v, want %v", r, want)
	}
}

func TestD2p2(t *testing.T) {
	want := 71220
	if r := d2p2(aocutils.ReadInput("d2p2.txt")); r != want {
		t.Errorf("d2p1(..) got %v, want %v", r, want)
	}
}
