package inputreader

import (
	"testing"
)

func TestAsString(t *testing.T) {
	if testStr := FirstString("testdata/textfile.txt"); testStr != "test" {
		t.Errorf("String expected to be 'test', got '%s'", testStr)
	}

	if testStr := FirstString("testdata/emptyfile.txt"); testStr != "" {
		t.Errorf("String expected to be '', got '%s'", testStr)
	}
}

func TestAsLines(t *testing.T) {
	expected := [...]string{"test", "another line", "empty line to come!"}
	lines := Lines("testdata/textfile.txt")

	if len(lines) != 3 {
		t.Errorf("Number of lines for test file expected to be 3, got %d", len(lines))
	}

	for i := range lines {
		if lines[i] != expected[i] {
			t.Errorf("Line %d expected to be '%s', got '%s'", i, expected[i], lines[i])
		}
	}

	emptyLines := Lines("testdata/emptyfile.txt")

	if len(emptyLines) != 0 {
		t.Errorf("Number of lines for empty file expected to be 0, got %d", len(emptyLines))
	}
}
