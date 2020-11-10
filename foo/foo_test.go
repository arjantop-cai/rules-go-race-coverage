package foo

import "testing"

func TestFoo(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Errorf("...")
	}
}
