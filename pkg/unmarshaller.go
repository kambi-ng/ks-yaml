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

	// context
	inlineComment    string
	hasInlineComment bool
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

// TODO write brackets here
func (m *unmarshaller) unmarshallNode(n ast.Node, depth int) {
	switch v := n.(type) {
	case *ast.BoolNode, *ast.FloatNode, *ast.IntegerNode, *ast.NullNode, *ast.StringNode:
		m.unmarshallValue(v, depth)
	case *ast.MappingValueNode:
		m.unmarshallKeyValueObj(v, depth)
	case *ast.MappingNode:
		m.unmarshallObject(v, depth)
	case *ast.SequenceNode:
		m.unmarshallArray(v, depth)
	default:
		fmt.Fprintf(&m.sb, "[x](%T)%s", n, v)
	}
}

func (m *unmarshaller) unmarshallKey(k ast.Node, depth int) {
	fmt.Fprintf(&m.sb, "%s: ", k.GetToken().Value)
	comm := k.GetComment()
	if comm != nil {
		ic := comm.GetToken().Value
		m.inlineComment = ic
		m.hasInlineComment = true
	}
}

func (m *unmarshaller) unmarshallObject(o *ast.MappingNode, depth int) {
	pre := strings.Repeat(m.indentString, depth)
	if depth != 0 {
		m.sb.WriteString("{")
	}

	if m.hasInlineComment {
		fmt.Fprintf(&m.sb, "%s#%s\n", pre, m.inlineComment)
	} else if depth != 0 {
		m.sb.WriteString("\n")
	}

	c := o.GetComment()
	if c != nil {
		fmt.Fprintf(&m.sb, "%s#%s\n", pre, c.GetToken().Value)
	}
	kvs := o.Values
	for i, kv := range kvs {

		kvc := kv.GetComment()
		if kvc != nil {
			fmt.Fprintf(&m.sb, " #%s\n", kvc.GetToken().Value)
		}

		m.sb.WriteString(pre)

		k := kv.Key
		v := kv.Value
		m.unmarshallKey(k, depth)
		m.unmarshallNode(v, depth+1)

		if i != len(kvs)-1 && depth != 0 {
			m.sb.WriteString(",")
		}
		m.sb.WriteString("\n")
	}
	if depth != 0 {
		pre := strings.Repeat(m.indentString, max(0, depth-1))
		fmt.Fprintf(&m.sb, "%s}", pre)
	}
}

func (m *unmarshaller) unmarshallKeyValueObj(n *ast.MappingValueNode, depth int) {
	pre := strings.Repeat(m.indentString, depth)
	if depth != 0 {
		m.sb.WriteString("{")
	}
	if m.hasInlineComment {
		fmt.Fprintf(&m.sb, " #%s\n", m.inlineComment)
	} else if depth != 0 {
		m.sb.WriteString("\n")
	}
	c := n.GetComment()
	if c != nil {
		fmt.Fprintf(&m.sb, "%s#%s\n", pre, c.GetToken().Value)
	}

	k := n.Key
	v := n.Value
	m.sb.WriteString(pre)
	m.unmarshallKey(k, depth)
	m.unmarshallNode(v, depth+1)
	if depth != 0 {
		pre := strings.Repeat(m.indentString, max(0, depth-1))
		fmt.Fprintf(&m.sb, "\n%s}", pre)
	}
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
		ic := v.GetComment().GetToken().Value
		m.inlineComment = ic
		m.hasInlineComment = true
	}
}

func (m *unmarshaller) unmarshallArray(n *ast.SequenceNode, depth int) {
	pre := strings.Repeat(m.indentString, depth)

	m.sb.WriteString("[")

	if m.hasInlineComment {
		fmt.Fprintf(&m.sb, "%s#%s\n", pre, m.inlineComment)
	} else if depth != 0 {
		m.sb.WriteString("\n")
	}

	c := n.GetComment()
	if c != nil {
		fmt.Fprintf(&m.sb, "%s#%s\n", pre, c.GetToken().Value)
	}

	v := n.Values
	for i, vv := range v {
		m.sb.WriteString(pre)
		m.unmarshallNode(vv, depth+1)
		if i != len(v)-1 && depth != 0 {
			m.sb.WriteString(",")
		}

		if m.hasInlineComment {
			fmt.Fprintf(&m.sb, " #%s", m.inlineComment)
			m.hasInlineComment = false
		}

		m.sb.WriteString("\n")
	}
	prec := strings.Repeat(m.indentString, max(0, depth-1))
	fmt.Fprintf(&m.sb, "%s]", prec)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
