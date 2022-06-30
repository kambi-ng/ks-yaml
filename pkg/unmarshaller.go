package ksyaml

import (
	"github.com/goccy/go-yaml/ast"
	"github.com/goccy/go-yaml/parser"
	"github.com/goccy/go-yaml/token"

	"fmt"
	"strings"
)

type unmarshaller struct {
	indentString string
	sb           strings.Builder
}

func newUnmarshaller(indentString string) *unmarshaller {
	return &unmarshaller{
		indentString: indentString,
		sb:           strings.Builder{},
	}
}

func (m *unmarshaller) unmarshallString(in string) (string, error) {
	inBytes := []byte(in)
	return m.unmarshallBytes(inBytes)
}

func (m *unmarshaller) unmarshallBytes(in []byte) (string, error) {
	f, err := parser.ParseBytes(in, parser.ParseComments)

	if err != nil {
		return "", err
	}

	for _, d := range f.Docs {
		docB := d.Body

		switch docB.(type) {
		case *ast.MappingNode:
			mnode := docB.(*ast.MappingNode)
			m.unmarshallMappingNode(mnode, 0)

		case *ast.MappingValueNode:
			mvnode := docB.(*ast.MappingValueNode)
			m.unmarshallMappingValue(mvnode, 0)
		}
	}

	return m.sb.String(), nil
}

func (m *unmarshaller) unmarshallMappingNode(data *ast.MappingNode, depth int) {
	pre := strings.Repeat(m.indentString, depth)

	comm := data.GetComment()
	if comm != nil {
		s := fmt.Sprintf("%s%s\n", pre, comm)
		m.sb.WriteString(s)
	}

	for _, val := range data.Values {
		m.unmarshallMappingValue(val, depth)
	}
}

func (m *unmarshaller) unmarshallMappingValue(data *ast.MappingValueNode, depth int) {
	pre := strings.Repeat(m.indentString, depth)

	comm := data.GetComment()
	if data.GetComment() != nil {
		s := fmt.Sprintf("%s%s\n", pre, comm)
		m.sb.WriteString(s)
	}

	key := data.Key
	val := data.Value

	switch val.(type) {
	case *ast.IntegerNode, *ast.FloatNode, *ast.BoolNode:
		s := fmt.Sprintf("%s%s: %s", pre, key, val)
		m.sb.WriteString(s)
	case *ast.StringNode:
		val := val.(*ast.StringNode)
		val.Token.Type = token.DoubleQuoteType
		fmt.Fprintf(&m.sb, "%s%s: %s", pre, key, val)
	case *ast.MappingNode:
		fmt.Fprintf(&m.sb, "%s%s: {\n", pre, key)
		m.unmarshallMappingNode(val.(*ast.MappingNode), depth+1)
		m.sb.WriteString(pre + "}")
	case *ast.MappingValueNode:
		fmt.Fprintf(&m.sb, "%s%s: {\n", pre, key)
		m.unmarshallMappingValue(val.(*ast.MappingValueNode), depth+1)
		m.sb.WriteString(pre + "}")
	case *ast.SequenceNode:
		fmt.Fprintf(&m.sb, "%s%s: [\n", pre, key)
		m.unmarshallSequenceNode(val.(*ast.SequenceNode), depth+1)
		m.sb.WriteString(pre + "]")
	default:
		s := fmt.Sprintf("[x] %s %s: %T", pre, key, val)
		m.sb.WriteString(s)
	}

	if depth > 0 {
		m.sb.WriteString(",")
	}
	m.sb.WriteString("\n")
}

func (m *unmarshaller) unmarshallSequenceNode(data *ast.SequenceNode, depth int) {
	pre := strings.Repeat(m.indentString, depth)

	comm := data.GetComment()
	if comm != nil {
		fmt.Fprintf(&m.sb, "%s%s\n", pre, comm)
	}

	for i, val := range data.Values {
		switch val.(type) {
		case *ast.IntegerNode, *ast.FloatNode, *ast.BoolNode:
			fmt.Fprintf(&m.sb, "%s%s", pre, val)
		case *ast.StringNode:
			val := val.(*ast.StringNode)
			val.Token.Type = token.DoubleQuoteType
			fmt.Fprintf(&m.sb, "%s%s", pre, val)
		case *ast.MappingNode:
			m.sb.WriteString(pre + "{\n")
			m.unmarshallMappingNode(val.(*ast.MappingNode), depth+1)
			m.sb.WriteString(pre + "}")
		case *ast.MappingValueNode:
			m.sb.WriteString(pre + "{\n")
			m.unmarshallMappingValue(val.(*ast.MappingValueNode), depth)
			m.sb.WriteString(pre + "}")
		case *ast.SequenceNode:
			m.sb.WriteString(pre + "[\n")
			m.unmarshallSequenceNode(val.(*ast.SequenceNode), depth+1)
			m.sb.WriteString(pre + "]")

		}
		if i != len(data.Values)-1 {
			m.sb.WriteString(",")
		}
		m.sb.WriteString("\n")
	}

}
