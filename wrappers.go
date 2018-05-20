package gamcrest

func Is(matcher Matcher) Matcher {
	return &PrefixWrapperMatcher{
		Prefix:   "is ",
		Delegate: matcher,
	}
}

type PrefixWrapperMatcher struct {
	Prefix   string
	Delegate Matcher
}

func (instance PrefixWrapperMatcher) Test(actual interface{}) (MatchResult, error) {
	delegateResult, err := instance.Delegate.Test(actual)
	if err != nil {
		return nil, err
	}
	return &PrefixWrapperMatchResult{
		Prefix:   instance.Prefix,
		Delegate: delegateResult,
	}, nil
}

type PrefixWrapperMatchResult struct {
	Prefix   string
	Delegate MatchResult
}

func (instance PrefixWrapperMatchResult) IsMatching() bool {
	return instance.Delegate.IsMatching()
}

func (instance PrefixWrapperMatchResult) IsFatal() bool {
	return instance.Delegate.IsMatching()
}

func (instance PrefixWrapperMatchResult) Describe(to DescriptionBuilder) error {
	to.AppendMessage(instance.Prefix)
	return instance.Delegate.Describe(to)
}

func (instance PrefixWrapperMatchResult) DescribeMismatch(to DescriptionBuilder) error {
	return instance.Delegate.DescribeMismatch(to)
}
