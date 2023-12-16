package day11

import (
	"testing"

	"github.com/agrigoryan/aoc_2023_go/aocutils"
)

func TestD11p1(t *testing.T) {
	want := 9608724
	if r := d11p1(aocutils.ReadInput("d11p1.txt")); r != want {
		t.Errorf("d11p1(..) got %v, want %v", r, want)
	}
}

func TestD11p2(t *testing.T) {
	want := 904633799472
	if r := d11p2(aocutils.ReadInput("d11p2.txt")); r != want {
		t.Errorf("d11p2(..) got %v, want %v", r, want)
	}
}
