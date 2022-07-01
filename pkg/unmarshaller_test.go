package ksyaml

import (
	"strings"
	"testing"
)

func TestConvertCommentNodeShouldReturnCommentString(t *testing.T) {
	expected_output := `# Test`

	input_string := `# Test`

	converter := NewConverter()
	output_string, _ := converter.Convert(input_string)

	if output_string != expected_output {
		t.Errorf("expected %s, got %s", expected_output, output_string)
	}
}

func TestConvertIntegerNodeShouldReturnInteger(t *testing.T) {
	input_string := `1`
	expected_output := `1`

	converter := NewConverter()
	output_string, _ := converter.Convert(input_string)

	if output_string != expected_output {
		t.Errorf("expected %s, got %s", expected_output, output_string)
	}
}

func TestConvertQuotedStringNodeShouldReturnQuotedString(t *testing.T) {
	input_string := `"test"`
	expected_output := `"test"`

	converter := NewConverter()
	output_string, _ := converter.Convert(input_string)

	if output_string != expected_output {
		t.Errorf("expected %s, got %s", expected_output, output_string)
	}
}

func TestConvertUnquotedStringNodeShouldReturnQuotedString(t *testing.T) {
	input_string := `test`
	expected_output := `"test"`

	converter := NewConverter()
	output_string, _ := converter.Convert(input_string)

	if output_string != expected_output {
		t.Errorf("expected %s, got %s", expected_output, output_string)
	}
}

func TestConvertBooleanNodeShouldReturnBoolean(t *testing.T) {
	input_string := `true`
	expected_output := `true`

	converter := NewConverter()
	output_string, _ := converter.Convert(input_string)

	if output_string != expected_output {
		t.Errorf("expected %s, got %s", expected_output, output_string)
	}
}

func TestConvertFloatNodeShouldReturnFloat(t *testing.T) {
	input_string := "3.1428"
	expected_output := "3.1428"

	converter := NewConverter()
	output_string, _ := converter.Convert(input_string)

	if output_string != expected_output {
		t.Errorf("expected %s, got %s", expected_output, output_string)
	}
}

func TestConvertTopLevelKeyPrimitivePairShouldReturnKeyPrimitive(t *testing.T) {
	input_string := `key-1: a`
	expected_output := `key-1: "a"`

	converter := NewConverter()
	output_string, _ := converter.Convert(input_string)

	if output_string != expected_output {
		t.Errorf("expected %s, got %s", expected_output, output_string)
	}
}

func TestConvertTopLevelKeyPrimitivePairWithCommentShouldReturnStringWithComment(t *testing.T) {
	input_string := `key-1: a # I AM A COMMENT`
	expected_output := `key-1: "a" # I AM A COMMENT`

	converter := NewConverter()
	output_string, _ := converter.Convert(input_string)

	if output_string != expected_output {
		t.Errorf("expected %s, got %s", expected_output, output_string)
	}
}

func TestConvertPrimitiveNodeWithCommentShouldReturnResultWithInlineComment(t *testing.T) {
	input_string := `i am a string # i am a comment`
	expected_output := `"i am a string" # i am a comment`

	converter := NewConverter()
	output_string, _ := converter.Convert(input_string)

	if output_string != expected_output {
		t.Errorf("expected %s, got %s", expected_output, output_string)
	}
}

func TestConvertMappingNode(t *testing.T) {
	input_string := `map:
 hello:
  hello:
   a: 4
   hello:
    hello: 2`
	expected_output := `map: {
 hello: {
  hello: {
   a: 4,
   hello: {
    hello: 2,
   },
  },
 },
}`

	converter := NewConverter()
	output_string, _ := converter.Convert(input_string)
	if strings.Trim(output_string, "\n") != strings.Trim(expected_output, "\n") {
		t.Errorf("expected \n%s\ngot\n%s", expected_output, output_string)
	}
}
