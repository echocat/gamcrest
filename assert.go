package gamcrest

import "testing"

func AssertThat(t *testing.T, actual interface{}, matcher Matcher) {
	result, err := matcher.Test(actual)
	if err != nil {
		t.Fatalf("Cannot execute %v: %v", matcher, err)
	}
	if !result.IsMatching() {
		db := BuildDescription()

		db.AppendMessage("\nExpected: ")
		if err := result.Describe(db); err != nil {
			t.Fatalf("Cannot describe of %v: %v", result, err)
		}

		db.AppendMessage("\n  Actual: ")
		if err := result.DescribeMismatch(db); err != nil {
			t.Fatalf("Cannot describe mismatch of %v: %v", result, err)
		}

		if result.IsFatal() {
			t.Fatal(db.String())
		} else {
			t.Error(db.String())
		}
	}
}
