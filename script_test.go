package dsl_test

import (
	"dsl-builder"
	"testing"
)

// TestRun calls dsl.Run, checking to ensure no error is returned.
func TestRun(t *testing.T) {
	var want error = nil
	err := dsl.Run("fakeScript", ";")

	if err != want {
		t.Fatalf("%v does not equal %v", err, want)
	}
}
