package day8

import (
	"testing"

	"github.com/agrigoryan/aoc_2023_go/aocutils"
)

func TestD8p1(t *testing.T) {
	want := 13301
	if r := d8p1(aocutils.ReadInput("d8p1.txt")); r != want {
		t.Errorf("d8p1(..) got %v, want %v", r, want)
	}
}

func TestD8p2(t *testing.T) {
	want := 6
	if r := d8p2(aocutils.ReadInput("d8p2.txt")); r != want {
		t.Errorf("d8p1(..) got %v, want %v", r, want)
	}
}
