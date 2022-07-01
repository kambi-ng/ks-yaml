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

func (m *unmarshaller) printInlineComment(node ast.Node, depth int) {
	comment := node.GetComment()
	if comment != nil {
		fmt.Fprintf(&m.sb, " %s", comment)
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

	for _, document := range f.Docs {
		documentBody := document.Body
		m.unmarshallNode(documentBody, 0)
	}

	return m.sb.String(), nil
}

func (m *unmarshaller) unmarshallPrimitiveNode(value ast.Node, depth int) {
	val := value.GetToken().Value
	switch value.(type) {
	case *ast.StringNode:
		fmt.Fprintf(&m.sb, "\"%s\"", val)

	case *ast.CommentNode, *ast.CommentGroupNode:
		fmt.Fprintf(&m.sb, "#%s", val)
	default:
		fmt.Fprintf(&m.sb, "%s", val)
	}
}

func (m *unmarshaller) unmarshallMappingNode(data *ast.MappingNode, depth int) {
	comment := data.GetComment()
	if comment != nil {
		fmt.Fprintf(&m.sb, "%s", comment)
		m.sb.WriteString("\n")
	}
	for i, val := range data.Values {
		m.unmarshallNode(val, depth)

		if i != len(data.Values)-1 {
			fmt.Fprint(&m.sb, ",")
			fmt.Fprintln(&m.sb)
		}
	}
}

func (m *unmarshaller) unmarshallInlineNode(key, value ast.Node, depth int) {
	pre := strings.Repeat(m.indentString, depth)
	m.sb.WriteString(pre)

	if key != nil {
		fmt.Fprintf(&m.sb, "%s: ", key.GetToken().Value)
	}

	switch value.(type) {
	case *ast.MappingNode, *ast.MappingValueNode:
		m.sb.WriteString("{")
		m.printInlineComment(key, depth)
		m.sb.WriteString("\n")
		m.unmarshallNode(value, depth+1)
		fmt.Fprintf(&m.sb, "\n%s}", pre)
	case *ast.SequenceNode:
		m.unmarshallSequenceNode(value.(*ast.SequenceNode), depth)
	default:
		m.unmarshallPrimitiveNode(value, depth)
	}
}

func (m *unmarshaller) unmarshallSequenceNode(node *ast.SequenceNode, depth int) {

	pre := strings.Repeat(m.indentString, depth)
	fmt.Fprintf(&m.sb, "[")
	m.printInlineComment(node, depth)
	m.sb.WriteString("\n")

	for i, val := range node.Values {
		m.unmarshallNode(val, depth+1)
		if i != len(node.Values)-1 {
			m.sb.WriteString(",")
		}
		m.printInlineComment(val, depth)
		if i != len(node.Values)-1 {
			m.sb.WriteString("\n")
		}
	}
	fmt.Fprintf(&m.sb, "\n%s]", pre)
}

func (m *unmarshaller) unmarshallNode(node ast.Node, depth int) {
	switch node.(type) {
	case *ast.IntegerNode, *ast.FloatNode, *ast.BoolNode, *ast.StringNode, *ast.CommentGroupNode, *ast.CommentNode:
		m.unmarshallInlineNode(nil, node, depth)
	case *ast.MappingValueNode:
		mappingValueNode := node.(*ast.MappingValueNode)
		m.unmarshallInlineNode(mappingValueNode.Key, mappingValueNode.Value, depth)
	case *ast.MappingNode:
		m.unmarshallMappingNode(node.(*ast.MappingNode), depth)
	case *ast.SequenceNode:
		m.unmarshallSequenceNode(node.(*ast.SequenceNode), depth)
	}
}
