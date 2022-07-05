package ksyaml

import (
	"github.com/goccy/go-yaml/parser"

	"fmt"
	"github.com/goccy/go-yaml/ast"
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
		m.unmarshallNode(docB, 0)
	}

	return m.sb.String(), nil
}

func (m *unmarshaller) unmarshallNode(n ast.Node, depth int) {
	switch v := n.(type) {
	case *ast.BoolNode, *ast.FloatNode, *ast.IntegerNode, *ast.NullNode, *ast.StringNode:
		m.unmarshallValue(v, depth)
	case *ast.MappingValueNode:
		m.unmarshallKeyValue(v, depth+1)
	case *ast.MappingNode:
		m.unmarshallObject(v, depth)
	case *ast.SequenceNode:
		m.unmarshallArray(v, depth)
	default:
		fmt.Fprintf(&m.sb, "[x](%T)%s", n, n)
	}
}

func (m *unmarshaller) unmarshallKey(k ast.Node, depth int) {
	pre := strings.Repeat(m.indentString, depth)
	m.sb.WriteString(pre)
	m.sb.WriteString(k.GetToken().Value)
	m.sb.WriteString(": ")

	comm := k.GetComment()
	if comm != nil {
		fmt.Fprintf(&m.sb, " #%s", comm.GetToken().Value)
	}
}

func (m *unmarshaller) unmarshallObject(o *ast.MappingNode, depth int) {

	if depth != 0 {
		m.sb.WriteString("{\n")
	}
	v := o.Values
	for i, kv := range v {
		m.unmarshallKeyValue(kv, depth)
		if i != len(v)-1 && depth != 0 {
			m.sb.WriteString(",")
		}
		m.sb.WriteString("\n")
	}
	if depth != 0 {
		m.sb.WriteString("}")
	}
}

func (m *unmarshaller) unmarshallKeyValue(n *ast.MappingValueNode, depth int) {
	k := n.Key
	v := n.Value
	m.unmarshallKey(k, depth)
	m.unmarshallNode(v, depth+1)
}

func (m *unmarshaller) unmarshallValue(v ast.Node, depth int) {

	vs := v.GetToken().Value
	switch v.(type) {
	case *ast.StringNode:
		fmt.Fprintf(&m.sb, `"%s"`, vs)
	default:
		m.sb.WriteString(vs)
	}
	if v.GetComment() != nil {
		fmt.Fprintf(&m.sb, " #%s", v.GetComment().GetToken().Value)
	}
}

func (m *unmarshaller) unmarshallArray(n *ast.SequenceNode, depth int) {
	m.sb.WriteString("[\n")
	v := n.Values
	for i, vv := range v {
		m.unmarshallNode(vv, depth+1)
		if i != len(v)-1 {
			m.sb.WriteString(",")
		}
		m.sb.WriteString("\n")
	}
	m.sb.WriteString("]\n")
}
