package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/goccy/go-yaml/ast"
	"github.com/goccy/go-yaml/parser"
	"github.com/goccy/go-yaml/token"
)

func main() {
	b, err := ioutil.ReadFile("test.yml")
	if err != nil {
		log.Fatal(err)
	}
	// f, _ := parser.ParseBytes(b, 0)
	f, _ := parser.ParseBytes(b, parser.ParseComments)
	var s string
	for _, v := range f.Docs {
		n := v.Body
		switch n.(type) {
		case *ast.MappingNode:
			n := n.(*ast.MappingNode)
			s += unmarshal(n, 0)
		}
	}
	// fmt.Println(f)
	fmt.Println(s)
}

func unmarshal(data *ast.MappingNode, depth int) string {
	var s string
	pre := strings.Repeat("  ", depth)
	c := data.GetComment()
	if c != nil {
		s += fmt.Sprintf("%s%s\n", pre, c)
	}
	for _, v := range data.Values {
		s += unmarshalMappingValue(v, depth)
	}
	return s
}

func unmarshalMappingValue(data *ast.MappingValueNode, depth int) string {
	var s string
	pre := strings.Repeat("  ", depth)
	n := data.Value
	c := data.GetComment()
	if c != nil {
		s += fmt.Sprintf("%s%s\n", pre, c)
	}
	switch n.(type) {
	case *ast.IntegerNode, *ast.FloatNode, *ast.BoolNode:
		s += fmt.Sprintf("%s%s: %s", pre, data.Key, n)
	case *ast.StringNode:
		n := n.(*ast.StringNode)
		n.Token.Type = token.DoubleQuoteType
		s += fmt.Sprintf("%s%s: %s", pre, data.Key, n)
	case *ast.MappingNode:
		k := data.Key.(*ast.StringNode).Value
		s += fmt.Sprintf("%s%s: {", pre, k)
		c := data.Key.GetComment()
		if c != nil {
			s += fmt.Sprintf(" %s\n", c)
		}
		n := n.(*ast.MappingNode)
		s += unmarshal(n, depth+1)
		s += fmt.Sprintf("%s}", pre)
	case *ast.SequenceNode:
		n := n.(*ast.SequenceNode)
		s += fmt.Sprintf("%s%s: [\n", pre, data.Key)
		s += unmarshalSequence(n, depth+1)
		s += fmt.Sprintf("%s]", pre)
	case *ast.MappingValueNode:
		n := n.(*ast.MappingValueNode)
		s += fmt.Sprintf("%s%s: {\n", pre, data.Key)
		s += unmarshalMappingValue(n, depth+1)
		s += fmt.Sprintf("%s}", pre)
	default:
		s += fmt.Sprintf("[x]%s%s: %T", pre, data.Key, n)
	}
	if depth != 0 {
		s += ","
	}
	s += "\n"
	return s
}

func unmarshalSequence(data *ast.SequenceNode, depth int) string {
	var s string
	pre := strings.Repeat("  ", depth)
	for i, v := range data.Values {
		switch v.(type) {
		case *ast.IntegerNode, *ast.FloatNode, *ast.BoolNode:
			s += fmt.Sprintf("%s%s", pre, v)
		case *ast.StringNode:
			n := v.(*ast.StringNode)
			n.Token.Type = token.DoubleQuoteType
			s += fmt.Sprintf("%s%s", pre, n)
		case *ast.MappingNode:
			n := v.(*ast.MappingNode)
			s += fmt.Sprintf("%s{", pre)
			s += unmarshal(n, depth+1)
			s += fmt.Sprintf("%s}", pre)
		case *ast.SequenceNode:
			n := v.(*ast.SequenceNode)
			s += fmt.Sprintf("%s[", pre)
			s += unmarshalSequence(n, depth+1)
			s += fmt.Sprintf("%s]", pre)
		case *ast.MappingValueNode:
			n := v.(*ast.MappingValueNode)
			s += unmarshalMappingValue(n, depth)
		}
		if i != len(data.Values)-1 {
			s += ","
		}
		s += "\n"
	}
	return s
}
