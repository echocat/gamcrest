package gamcrest

type Matcher interface {
	Test(actual interface{}) (MatchResult, error)
}

type MatchResult interface {
	IsMatching() bool
	IsFatal() bool
	Describe(to DescriptionBuilder) error
	DescribeMismatch(to DescriptionBuilder) error
}

type StaticMatchResult struct {
	Value               bool
	Fatal               bool
	Description         DescriptionBuilder
	MismatchDescription DescriptionBuilder
}

func (instance StaticMatchResult) IsMatching() bool {
	return instance.Value
}

func (instance StaticMatchResult) IsFatal() bool {
	return instance.Fatal
}

func (instance StaticMatchResult) Describe(to DescriptionBuilder) error {
	to.Append(instance.Description)
	return nil
}

func (instance StaticMatchResult) DescribeMismatch(to DescriptionBuilder) error {
	to.Append(instance.MismatchDescription)
	return nil
}
