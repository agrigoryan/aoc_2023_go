package day3

import (
	"testing"

	"github.com/agrigoryan/aoc_2023_go/aocutils"
)

func TestD3p1(t *testing.T) {
	want := 553825
	if r := d3p1(aocutils.ReadInput("d3p1.txt")); r != want {
		t.Errorf("d3p1(..) got %v, want %v", r, want)
	}
}

func TestD3p2(t *testing.T) {
	want := 93994191
	if r := d3p2(aocutils.ReadInput("d3p2.txt")); r != want {
		t.Errorf("d3p1(..) got %v, want %v", r, want)
	}
}
