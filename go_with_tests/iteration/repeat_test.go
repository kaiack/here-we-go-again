package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	n := 10
	for b.Loop() {
		Repeat("a", n)
	}
}

func ExampleRepeat() {
	repeated := Repeat("6", 7)
	fmt.Println(repeated)
	// Output: 6666666
}
