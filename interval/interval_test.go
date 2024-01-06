package interval

import "testing"

func TestCorrectIntervalConstruction(t *testing.T) {
	i, _ := New(1, 5)
	if i.Start != 1 {
		t.Errorf("Expected start to be 1, got %d", i.Start)
	}

	if i.End != 5 {
		t.Errorf("Expected end to be 5, got %d", i.End)
	}
}

func TestInvertedIntervalConstruction(t *testing.T) {
	i, _ := New(5, 1)
	// expect an error
	if i != nil {
		t.Errorf("Expected error, got none")
	}

}
