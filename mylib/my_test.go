package mylib_test

import (
	"testing"

	"github.com/dandeliondeathray/goosedev/mylib"
)

func TestAdder_2Plus1_Is3(t *testing.T) {
	adder := mylib.NewAdder(2)
	result := adder.Op(1)
	if result != 3 {
		t.Fatalf("Expected 3 but got %d", result)
	}
}

func TestSubtractor_4Minus1_Is3(t *testing.T) {
	subtractor := mylib.NewSubtractor(4)
	result := subtractor.Op(1)
	if result != 3 {
		t.Fatalf("Expected 3 but got %d", result)
	}
}
