package day4

import (
	"testing"

	"github.com/agrigoryan/aoc_2023_go/aocutils"
)

func TestD4p1(t *testing.T) {
	want := 22193
	if r := d4p1(aocutils.ReadInput("d4p1.txt")); r != want {
		t.Errorf("d4p1(..) got %v, want %v", r, want)
	}
}

func TestD4p2(t *testing.T) {
	want := 5625994
	if r := d4p2(aocutils.ReadInput("d4p2.txt")); r != want {
		t.Errorf("d4p1(..) got %v, want %v", r, want)
	}
}
