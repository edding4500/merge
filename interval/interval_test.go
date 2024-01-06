package interval

import "testing"

// Test construction of an interval with start > end
func TestCorrectIntervalConstruction(t *testing.T) {
	i, _ := New(1, 5)
	if i.Start != 1 {
		t.Errorf("Expected start to be 1, got %d", i.Start)
	}

	if i.End != 5 {
		t.Errorf("Expected end to be 5, got %d", i.End)
	}
}

// Test construction of an interval with start < end
func TestInvertedIntervalConstruction(t *testing.T) {
	i, _ := New(5, 1)
	// expect an error
	if i != nil {
		t.Errorf("Expected error, got none")
	}

}

// Test interval construction with start == end
func TestIntervalConstructionWithEqualStartAndEnd(t *testing.T) {
	i, _ := New(1, 1)
	if i.Start != 1 {
		t.Errorf("Expected start to be 1, got %d", i.Start)
	}

	if i.End != 1 {
		t.Errorf("Expected end to be 1, got %d", i.End)
	}
}

// Test interval construction with negative numbers
func TestIntervalConstructionWithNegativeNumbers(t *testing.T) {
	i, error := New(-5, -1)
	if error != nil {
		t.Errorf("Expected no error, got %s", error)
	}
	if i.Start != -5 {
		t.Errorf("Expected start to be -5, got %d", i.Start)
	}

	if i.End != -1 {
		t.Errorf("Expected end to be -1, got %d", i.End)
	}
}

// Test interval construction with negative numbers
// where start > end
func TestIntervalConstructionWithNegativeNumbersInverted(t *testing.T) {
	_, error := New(-1, -5)
	if error == nil {
		t.Errorf("Expected error, got none")
	}
}
