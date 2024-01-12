package cal

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Error("Expected 3, got ", result)
	}
}
