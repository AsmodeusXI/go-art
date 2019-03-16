package buildstring

import (
	"testing"
)

func TestBuild(t *testing.T) {
	expected := "One Two Three"
	actual := Build("One", "Two", "Three")
	if actual != expected {
		t.Errorf("Build(%q, %q, %q) == %q, expected %q", "One", "Two", "Three", actual, expected)
	}
}
