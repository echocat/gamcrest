package gamcrest

import (
	"reflect"
)

func Equals(expected interface{}) *EqualsMatcher {
	return &EqualsMatcher{
		Expected:                       expected,
		ExpectedStateDescriptionPrefix: "equals ",
	}
}

func IsEqual(expected interface{}) *EqualsMatcher {
	return &EqualsMatcher{
		Expected:                       expected,
		ExpectedStateDescriptionPrefix: "is equal ",
	}
}

func IsEqualTo(expected interface{}) *EqualsMatcher {
	return &EqualsMatcher{
		Expected:                       expected,
		ExpectedStateDescriptionPrefix: "is equal to ",
	}
}

func EqualTo(expected interface{}) *EqualsMatcher {
	return &EqualsMatcher{
		Expected:                       expected,
		ExpectedStateDescriptionPrefix: "equal to ",
	}
}

type EqualsMatcher struct {
	Expected                       interface{}
	ExpectedStateDescriptionPrefix string
}

func (instance EqualsMatcher) Test(actual interface{}) (MatchResult, error) {
	return &StaticMatchResult{
		Value: instance.test(actual),
		Description: BuildDescription().
			AppendMessage(instance.ExpectedStateDescriptionPrefix).
			AppendValue(instance.Expected),
		MismatchDescription: BuildDescription().
			AppendMessage("but was ").
			AppendValue(actual),
	}, nil
}

func (instance EqualsMatcher) test(actual interface{}) bool {
	return reflect.DeepEqual(instance.Expected, actual)
}
