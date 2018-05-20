package gamcrest

import "testing"

func TestIsMatcher_Test(t *testing.T) {
	instance := Is(EqualTo("foo"))

	actual, err := instance.Test("bar")

	if err != nil {
		t.Fatal(err)
	}
	if actual.IsMatching() {
		t.Fatal(actual)
	}

	description := BuildDescription()
	if err := actual.Describe(description); err != nil {
		t.Fatal(err)
	}
	if description.String() != "is equal to <foo>" {
		t.Fatal(description)
	}

	mismatchDescription := BuildDescription()
	if err := actual.DescribeMismatch(mismatchDescription); err != nil {
		t.Fatal(err)
	}
	if mismatchDescription.String() != "but was <bar>" {
		t.Fatal(mismatchDescription)
	}
}