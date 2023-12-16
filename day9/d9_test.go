package day9

import (
	"testing"

	"github.com/agrigoryan/aoc_2023_go/aocutils"
)

func TestD9p1(t *testing.T) {
	want := 1887980197
	if r := d9p1(aocutils.ReadInput("d9p1.txt")); r != want {
		t.Errorf("d9p1(..) got %v, want %v", r, want)
	}
}

func TestD9p2(t *testing.T) {
	want := 2
	if r := d9p2(aocutils.ReadInput("d9p2.txt")); r != want {
		t.Errorf("d9p1(..) got %v, want %v", r, want)
	}
}
