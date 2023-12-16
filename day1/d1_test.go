package day1

import (
	"testing"

	"github.com/agrigoryan/aoc_2023_go/aocutils"
)

func TestD1p1(t *testing.T) {
	want := 55447
	if r := d1p1(aocutils.ReadInput("d1p1.txt")); r != want {
		t.Errorf("d1p1(..) got %v, want %v", r, want)
	}
}

func TestD1p2(t *testing.T) {
	want := 54706
	if r := d1p2(aocutils.ReadInput("d1p2.txt")); r != want {
		t.Errorf("d1p1(..) got %v, want %v", r, want)
	}
}
