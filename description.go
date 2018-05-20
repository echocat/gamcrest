package gamcrest

import (
	"strings"
	"reflect"
	"fmt"
)

type DescriptionBuilder interface {
	AppendMessage(string, ...interface{}) DescriptionBuilder
	AppendValue(interface{}) DescriptionBuilder
	Append(DescriptionBuilder) DescriptionBuilder
	String() string
}

func BuildDescription() DescriptionBuilder {
	return &DefaultDescriptionBuilder{}
}

type DefaultDescriptionBuilder struct {
	SB strings.Builder
}

func (instance DefaultDescriptionBuilder) String() string {
	return instance.SB.String()
}

func (instance *DefaultDescriptionBuilder) Append(other DescriptionBuilder) DescriptionBuilder {
	return instance.AppendMessage(other.String())
}

func (instance *DefaultDescriptionBuilder) AppendMessage(message string, args ... interface{}) DescriptionBuilder {
	return instance.appendMessage("", message, args...)
}

func (instance *DefaultDescriptionBuilder) appendMessage(indent string, message string, args ... interface{}) *DefaultDescriptionBuilder {
	if _, err := instance.SB.WriteString(indent); err != nil {
		panic(err)
	}
	content := message
	if len(args) > 0 {
		content = fmt.Sprintf(message, args...)
	}
	if _, err := instance.SB.WriteString(content); err != nil {
		panic(err)
	}
	return instance
}

func (instance *DefaultDescriptionBuilder) AppendValue(value interface{}) DescriptionBuilder {
	return instance.appendValue("", "", reflect.ValueOf(value))
}

func (instance *DefaultDescriptionBuilder) appendValue(indent string, prefix string, value reflect.Value) *DefaultDescriptionBuilder {
	if method, ok := value.Type().MethodByName("String"); ok {
		mt := method.Type
		if mt.NumOut() == 1 && mt.Out(0) == reflect.TypeOf("") {
			return instance.appendMessage(indent, "%v", value.Interface())
		}
	}

	switch value.Kind() {
	case reflect.Ptr:
		return instance.appendValue(indent, "*", value.Elem())
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		return instance.appendSlice(indent, prefix, value)
	case reflect.Map:
		return instance.appendMap(indent, prefix, value)
	case reflect.Struct:
		return instance.appendStruct(indent, prefix, value)
	case reflect.String:
		return instance.appendMessage(indent, "%s<%s>", prefix, value.Interface())
	}
	return instance.appendMessage(indent, "%s%v", prefix, value.Interface())
}

func (instance *DefaultDescriptionBuilder) appendSlice(indent string, prefix string, value reflect.Value) *DefaultDescriptionBuilder {
	instance.appendMessage(indent, "%s%v{\n", prefix, value.Type())
	l := value.Len()
	for i := 0; i < l; i++ {
		instance.
			appendValue(indent+"\t", "", value.Index(i)).
			appendMessage("", ",\n")
	}
	return instance.appendRune('}')
}

func (instance *DefaultDescriptionBuilder) appendMap(indent string, prefix string, value reflect.Value) *DefaultDescriptionBuilder {
	instance.appendMessage(indent, "%s%v{\n", prefix, value.Type())
	for _, key := range value.MapKeys() {
		instance.
			appendValue(indent+"\t", "", key).
			appendMessage("", ": ").
			appendValue("", "", value.MapIndex(key)).
			appendMessage("", ",\n")
	}
	return instance.appendRune('}')
}

func (instance *DefaultDescriptionBuilder) appendStruct(indent string, prefix string, value reflect.Value) *DefaultDescriptionBuilder {
	t := value.Type()
	l := value.NumField()
	instance.appendMessage(indent, "%s%v{\n", prefix, value.Type())
	for i := 0; i < l; i++ {
		field := t.Field(i)
		instance.
			appendMessage(indent+"\t", "", field.Name).
			appendMessage("", ": ").
			appendValue("", "", value.Field(i)).
			appendMessage("", ",\n")
	}
	return instance.appendRune('}')
}

func (instance *DefaultDescriptionBuilder) appendRune(r rune) *DefaultDescriptionBuilder {
	if _, err := instance.SB.WriteRune(r); err != nil {
		panic(err)
	}
	return instance
}
