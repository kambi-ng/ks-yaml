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
	in = append(in, byte('\n'))
	f, err := parser.ParseBytes(in, parser.ParseComments)

	if err != nil {
		return "", err
	}

	for _, d := range f.Docs {
		docB := d.Body
		m.unmarshallNode(docB, 0)
		m.writeInlineComment()
		m.sb.WriteString("\n")
	}

	return m.sb.String(), nil
}

func (m *unmarshaller) unmarshallNode(n ast.Node, depth int) {
	pre := strings.Repeat(m.indentString, max(0, depth-1))

	switch v := n.(type) {
	case *ast.BoolNode, *ast.FloatNode, *ast.IntegerNode, *ast.NullNode, *ast.StringNode:
		m.unmarshallValue(v, depth)
	case *ast.MappingValueNode:
		if depth != 0 {
			m.sb.WriteString("{")
			m.unmarshallKeyValueObj(v, depth)
			fmt.Fprintf(&m.sb, "\n%s}", pre)
			break
		}
		m.unmarshallKeyValueObj(v, depth)
	case *ast.MappingNode:
		if depth != 0 {
			m.sb.WriteString("{")
			m.unmarshallObject(v, depth)
			fmt.Fprintf(&m.sb, "%s}", pre)
			break
		}
		m.unmarshallObject(v, depth)
	case *ast.SequenceNode:
		m.sb.WriteString("[")
		m.unmarshallArray(v, depth)
		fmt.Fprintf(&m.sb, "%s]", pre)
	case *ast.LiteralNode:
		m.unmarshallLiteral(v, depth)
	case *ast.InfinityNode, *ast.NanNode:
		 m.unmarshallSpecialMathNode(v, depth)
	// TODO other nodes
	// Anchor node
	// alias node
	// comment node and comment group node
	// directive node
	// Merge Key node
	// Tag Node
	default:
		fmt.Fprintf(&m.sb, "[x](%T)%s", n, v)
	}
}

func (m *unmarshaller) unmarshallSpecialMathNode(n ast.Node, depth int) {
	m.sb.WriteString(n.GetToken().Value)
	if n.GetComment() != nil {
		ic := n.GetComment().GetToken().Value
		m.addInlineComment(ic)
	}
}

func (m *unmarshaller) unmarshallLiteral(n *ast.LiteralNode, depth int) {

	if depth <= 1 {
		fmt.Fprintf(&m.sb, "%s", n.String())
		return
	}

	origin := n.Value.GetToken().Origin
	lit := strings.TrimSpace(origin)

	fmt.Fprintf(&m.sb, `"%s"`, lit)
	if n.GetComment() != nil {
		c := n.GetComment().GetToken().Value
		m.addInlineComment(c)
	}
}

func (m *unmarshaller) unmarshallKey(k ast.Node, depth int) {
	fmt.Fprintf(&m.sb, "%s: ", k.GetToken().Value)
	comm := k.GetComment()
	if comm != nil {
		ic := comm.GetToken().Value
		m.addInlineComment(ic)
	}
}

func (m *unmarshaller) unmarshallObject(o *ast.MappingNode, depth int) {
	pre := strings.Repeat(m.indentString, depth)
	m.writeInlineComment()
	m.sb.WriteString("\n")
	c := o.GetComment()
	if c != nil {
		fmt.Fprintf(&m.sb, "%s#%s\n", pre, c.GetToken().Value)
	}
	kvs := o.Values
	for i, kv := range kvs {

		kvc := kv.GetComment()
		if kvc != nil {
			fmt.Fprintf(&m.sb, "%s#%s\n", pre, kvc.GetToken().Value)
		}

		m.sb.WriteString(pre)

		k := kv.Key
		v := kv.Value
		m.unmarshallKey(k, depth)
		m.unmarshallNode(v, depth+1)

		if i != len(kvs)-1 && depth != 0 {
			m.sb.WriteString(",")
		}
		m.writeInlineComment()
		m.sb.WriteString("\n")
	}
}

func (m *unmarshaller) unmarshallKeyValueObj(n *ast.MappingValueNode, depth int) {
	pre := strings.Repeat(m.indentString, depth)
	m.writeInlineComment()
	m.sb.WriteString("\n")

	c := n.GetComment()
	if c != nil {
		fmt.Fprintf(&m.sb, "%s#%s\n", pre, c.GetToken().Value)
	}

	k := n.Key
	v := n.Value
	m.sb.WriteString(pre)
	m.unmarshallKey(k, depth)
	m.unmarshallNode(v, depth+1)
	m.writeInlineComment()

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
		m.addInlineComment(ic)
	}
}

func (m *unmarshaller) unmarshallArray(n *ast.SequenceNode, depth int) {
	pre := strings.Repeat(m.indentString, depth)
	m.writeInlineComment()
	m.sb.WriteString("\n")

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
		m.writeInlineComment()
		m.sb.WriteString("\n")
	}
}

func (m *unmarshaller) writeInlineComment() {
	if !m.hasInlineComment {
		return
	}
	fmt.Fprintf(&m.sb, " #%s", m.inlineComment)
	m.hasInlineComment = false
}

func (m *unmarshaller) addInlineComment(ic string) {
	m.hasInlineComment = true
	m.inlineComment = ic
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
