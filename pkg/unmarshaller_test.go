package ksyaml

import (
	"strings"
	"testing"
)

func TestInlineNode(t *testing.T) {
	t.Run("Integer", func(t *testing.T) {
		in := `1`
		ex := `1`

		conv := NewConverter()
		out, _ := conv.Convert(in)

		if out != ex {
			t.Errorf("expected %s, got %s", ex, out)
		}
	})
	t.Run("Bool", func(t *testing.T) {
		in := `true`
		ex := `true`

		conv := NewConverter()
		out, _ := conv.Convert(in)

		if out != ex {
			t.Errorf("expected %s, got %s", ex, out)
		}
	})
	t.Run("Float", func(t *testing.T) {
		in := "3.1428"
		ex := "3.1428"

		conv := NewConverter()
		out, _ := conv.Convert(in)

		if out != ex {
			t.Errorf("expected %s, got %s", ex, out)
		}
	})
	t.Run("string", func(t *testing.T) {
		t.Run("single quoted", func(t *testing.T) {
			in := `'test'`
			ex := `"test"`

			conv := NewConverter()
			out, _ := conv.Convert(in)

			if out != ex {
				t.Errorf("expected %s, got %s", ex, out)
			}
		})

		t.Run("double quoted", func(t *testing.T) {
			in := `"test"`
			ex := `"test"`

			conv := NewConverter()
			out, _ := conv.Convert(in)

			if out != ex {
				t.Errorf("expected %s, got %s", ex, out)
			}
		})
		t.Run("unquoted", func(t *testing.T) {
			in := `test`
			ex := `"test"`

			conv := NewConverter()
			out, _ := conv.Convert(in)

			if out != ex {
				t.Errorf("expected %s, got %s", ex, out)
			}
		})
	})
	t.Run("mapping", func(t *testing.T) {
		in := `key-1: a`
		ex := `key-1: "a"`

		conv := NewConverter()
		out, _ := conv.Convert(in)

		if out != ex {
			t.Errorf("expected %s, got %s", ex, out)
		}
	})
}

func TestComment(t *testing.T) {
	t.Run("basic inline", func(t *testing.T) {
		ex := `# Test`
		in := `# Test`

		conv := NewConverter()
		out, _ := conv.Convert(in)

		if out != ex {
			t.Errorf("expected %s, got %s", ex, out)
		}
	})
	t.Run("top level on map", func(t *testing.T) {
		in := `key-1: a # I AM A COMMENT`
		ex := `key-1: "a" # I AM A COMMENT`

		conv := NewConverter()
		out, _ := conv.Convert(in)

		if out != ex {
			t.Errorf("expected %s, got %s", ex, out)
		}
	})

	t.Run("top level on primitive", func(t *testing.T) {
		in := `i am a string # i am a comment`
		ex := `"i am a string" # i am a comment`

		conv := NewConverter()
		out, _ := conv.Convert(in)

		if out != ex {
			t.Errorf("expected %s, got %s", ex, out)
		}
	})
	t.Run("mapping", func(t *testing.T) {
		in := `# HELLO
test:
  test:
    test-1: 4
    test:
      test: 3

hello:
  world: test`

		ex := `# HELLO
test: {
  test: {
    test-1: 4,
    test: {
      test: 3
    }
  }
}
hello: {
  world: "test"
}`

		conv := NewConverter()
		out, _ := conv.Convert(in)
		if strings.Trim(out, "\n") != strings.Trim(ex, "\n") {
			t.Errorf("expected \n%s\ngot\n%s", ex, out)
		}
	})
}

func TestConvertMappingNode(t *testing.T) {
	in := `map:
 hello:
  hello:
   a: 4
   hello:
    hello: 2`
	ex := `map: {
  hello: {
    hello: {
      a: 4,
      hello: {
        hello: 2
      }
    }
  }
}
`

	conv := NewConverter()
	out, _ := conv.Convert(in)
	if strings.TrimSpace(out) != strings.TrimSpace(ex) {
		t.Errorf("expected \n%s\ngot\n%s", ex, out)
	}
}
