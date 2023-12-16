package day6

import (
	"testing"

	"github.com/agrigoryan/aoc_2023_go/aocutils"
)

func TestD6p1(t *testing.T) {
	want := 588588
	if r := d6p1(aocutils.ReadInput("d6p1.txt")); r != want {
		t.Errorf("d6p1(..) got %v, want %v", r, want)
	}
}

func TestD6p2(t *testing.T) {
	want := 34655848
	if r := d6p2(aocutils.ReadInput("d6p2.txt")); r != want {
		t.Errorf("d6p1(..) got %v, want %v", r, want)
	}
}
