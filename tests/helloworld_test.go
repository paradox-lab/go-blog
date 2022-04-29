package tests

import (
	"fmt"
	"testing"
)

func Abs(int2 int) int {
	fmt.Println(int2)
	return 1
}

func TestAbs(t *testing.T) {
	got := Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}
