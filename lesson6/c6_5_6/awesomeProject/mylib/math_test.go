package mylib

import (
	"fmt"
	"testing"
)

var Debug bool = true

func Example() {
	v := Average([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(v)
}

func ExampleAverage() {
	v := Average([]int{1, 2, 3, 4, 5})
	fmt.Println(v)
}

func TestAverage(t *testing.T) {
	if Debug {
		t.Skip("Skip Reason")
	}
	v := Average([]int{1, 2, 3, 4, 5})
	if v != 3 {
		t.Error("Expected 3, got", v)
	}
}
