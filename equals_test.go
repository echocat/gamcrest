package gamcrest

import "testing"

func TestEqualsMatcher_Test(t *testing.T) {
	AssertThat(t, "foo", IsEqualTo("foo"))
	AssertThat(t, "foo", IsEqualTo("bar"))
}
