package ksyaml

import "strings"

type Converter struct {
	indentation int
	withTab     bool
}

type ConverterOption func(*Converter)

func InsertWhitespace(in string) string {
	input_string := strings.Builder{}

	if !strings.HasPrefix(in, "\n") {
		input_string.WriteString("\n")
	}
	input_string.WriteString(in)
	if !strings.HasSuffix(in, "\n") {
		input_string.WriteString("\n")
	}

	return input_string.String()
}

func WithIndentation(indentation int) ConverterOption {
	return func(c *Converter) {
		c.indentation = indentation
	}
}

func NewConverter(opts ...ConverterOption) *Converter {
	c := &Converter{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Converter) Convert(in string) (string, error) {
	s := " "

	if c.indentation == 0 {
		c.indentation = 2
	}
	indentString := strings.Repeat(s, c.indentation)

	m := newUnmarshaller(indentString)

	return m.unmarshallString(InsertWhitespace(in))
}

func Convert(in string) (string, error) {
	return NewConverter(
		WithIndentation(2),
	).Convert(in)
}
