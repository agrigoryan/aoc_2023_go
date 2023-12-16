package day10

import (
	"testing"

	"github.com/agrigoryan/aoc_2023_go/aocutils"
)

func TestD10p1(t *testing.T) {
	want := 6613
	if r := d10p1(aocutils.ReadInput("d10p1.txt")); r != want {
		t.Errorf("d10p1(..) got %v, want %v", r, want)
	}
}

func TestD10p2(t *testing.T) {
	want := 511
	if r := d10p2(aocutils.ReadInput("d10p2.txt")); r != want {
		t.Errorf("d10p1(..) got %v, want %v", r, want)
	}
}
