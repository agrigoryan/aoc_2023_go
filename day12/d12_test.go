package day12

import (
	"testing"

	"github.com/agrigoryan/aoc_2023_go/aocutils"
)

func TestD12p1(t *testing.T) {
	want := 7694
	if r := d12p1(aocutils.ReadInput("d12p1.txt")); r != want {
		t.Errorf("d12p1(..) got %v, want %v", r, want)
	}
}

func TestD12p2(t *testing.T) {
	want := 1
	if r := d12p2(aocutils.ReadInput("d12p2.txt")); r != want {
		t.Errorf("d12p2(..) got %v, want %v", r, want)
	}
}
