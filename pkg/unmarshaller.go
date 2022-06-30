package ksyaml

import (
	"github.com/goccy/go-yaml/ast"
	"github.com/goccy/go-yaml/parser"

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

		case *ast.IntegerNode, *ast.FloatNode, *ast.BoolNode, *ast.StringNode:
			commAfter := m.unmarshallInlineNode(nil, docB, 0)
			if commAfter != "" {
				fmt.Fprintf(&m.sb, " #%s", commAfter)
			}

		default:
			fmt.Fprintf(&m.sb,"[x] its  %T %s", docB, docB)
		}
	}

	return m.sb.String(), nil
}

func (m *unmarshaller) unmarshallMappingNode(data *ast.MappingNode, depth int) {
	pre := strings.Repeat(m.indentString, depth)

	comm := data.GetComment()
	if comm != nil {
		fmt.Fprintf(&m.sb, "%s%s\n", pre, comm)
	}

	for _, val := range data.Values {
		m.unmarshallMappingValue(val, depth)
	}
}

func (m *unmarshaller) unmarshallMappingValue(data *ast.MappingValueNode, depth int) {
	pre := strings.Repeat(m.indentString, depth)

	comm := data.GetComment()
	if data.GetComment() != nil {
		fmt.Fprintf(&m.sb,"%s%s\n", pre, comm)
	}

	key := data.Key
	val := data.Value

	commAfter := ""

	switch val.(type) {
	case *ast.IntegerNode, *ast.FloatNode, *ast.BoolNode, *ast.StringNode:
		commAfter = m.unmarshallInlineNode(key, val, depth)
	case *ast.MappingNode:
		fmt.Fprintf(&m.sb, "%s%s: {", pre, key.GetToken().Value)
		comm := key.GetComment()
		if comm != nil {
			fmt.Fprintf(&m.sb," %s", comm)
		}
		m.sb.WriteString("\n")
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
		fmt.Fprintf(&m.sb,"[x] %s %s: %T", pre, key, val)
	}

	if depth > 0 {
		m.sb.WriteString(",")
	}
	if commAfter != "" {
		fmt.Fprintf(&m.sb, " #%s", commAfter)
	}
	m.sb.WriteString("\n")
}

func (m *unmarshaller) unmarshallInlineNode(key, value ast.Node, depth int) string {
	pre := strings.Repeat(m.indentString, depth)

	m.sb.WriteString(pre)
	val := value.GetToken().Value
	if key != nil {
		fmt.Fprintf(&m.sb, "%s: ", key.GetToken().Value)
	}
	switch value.(type) {
	case *ast.StringNode:
		fmt.Fprintf(&m.sb, "\"%s\"", val)
	default:
		fmt.Fprintf(&m.sb, "%s", val)
	}
	if value.GetComment() != nil {
		return value.GetComment().GetToken().Value
	}
	return ""
}

func (m *unmarshaller) unmarshallSequenceNode(data *ast.SequenceNode, depth int) {
	pre := strings.Repeat(m.indentString, depth)

	comm := data.GetComment()
	if comm != nil {
		fmt.Fprintf(&m.sb, "%s%s\n", pre, comm)
	}

	commAfter := ""
	for i, val := range data.Values {
		switch val.(type) {
		case *ast.IntegerNode, *ast.FloatNode, *ast.BoolNode, *ast.StringNode:
			commAfter = m.unmarshallInlineNode(nil, val, depth)
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
		if commAfter != "" {
			fmt.Fprintf(&m.sb, " #%s", commAfter)
		}
		m.sb.WriteString("\n")
	}

}
