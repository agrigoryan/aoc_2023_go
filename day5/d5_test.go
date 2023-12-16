package day5

import (
	"testing"

	"github.com/agrigoryan/aoc_2023_go/aocutils"
)

func TestD5p1(t *testing.T) {
	want := 662197086
	if r := d5p1(aocutils.ReadInput("d5p1.txt")); r != want {
		t.Errorf("d5p1(..) got %v, want %v", r, want)
	}
}

func TestD5p2(t *testing.T) {
	want := 52510809
	if r := d5p2(aocutils.ReadInput("d5p2.txt")); r != want {
		t.Errorf("d5p1(..) got %v, want %v", r, want)
	}
}
