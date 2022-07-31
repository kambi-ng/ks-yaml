package ksyaml

import "strings"

type Converter struct {
	indentation int
	withTab     bool
}

type ConverterOption func(*Converter)

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
		c.indentation = 1
	}
    indentString := strings.Repeat(s, c.indentation)

	m := newUnmarshaller(indentString)

	return m.unmarshallString(in)
}

func Convert(in string) (string, error) {
	return NewConverter(
		WithIndentation(2),
	).Convert(in)
}
