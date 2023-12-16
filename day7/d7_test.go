package day7

import (
	"testing"

	"github.com/agrigoryan/aoc_2023_go/aocutils"
)

func TestD7p1(t *testing.T) {
	want := 250951660
	if r := d7p1(aocutils.ReadInput("d7p1.txt")); r != want {
		t.Errorf("d7p1(..) got %v, want %v", r, want)
	}
}

func TestD7p2(t *testing.T) {
	want := 251156055
	if r := d7p2(aocutils.ReadInput("d7p2.txt")); r != want {
		t.Errorf("d7p1(..) got %v, want %v", r, want)
	}
}
