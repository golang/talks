// +build OMIT

package sample // OMIT

func TestSample(t *testing.T) { // OMIT
	// Typical test code
	if got, want := testTargetFunc(input), expectedValue; !checkTestResult(got, want) {
		t.Errorf("testTargetFunc(%v) = %v; want %v", input, got, want)
	}
} // OMIT
