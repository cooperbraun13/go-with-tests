package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	// ... setup ...
	for b.Loop() {
		// ... code to measure ...
		Repeat("a", 6)
	}
	// ... cleanup ...
}

func ExampleRepeat() {
	result := Repeat("c", 3)
	fmt.Println(result)
	// Output: ccc
}
