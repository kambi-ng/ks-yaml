package ksyaml

import "testing"
import "strings"

const oneSpace = " "

func testEqual(t *testing.T, exp, out string) {
	exp = strings.TrimSpace(exp)
	out = strings.TrimSpace(out)
	if exp != out {
		t.Errorf("Expected:\n%s\n\nGot:\n\n%s", exp, out)
	}
}

func TestUnmarshallerSimple(t *testing.T) {
	t.Run("Simple string key value", func(t *testing.T) {
		um := newUnmarshaller(oneSpace)

		in := `key: value`
		exp := `key: "value"`
		out, err := um.unmarshallString(in)

		if err != nil {
			t.Errorf("Error: %s", err)
			return
		}

		testEqual(t, exp, out)
	})

	t.Run("Simple string key value with comment", func(t *testing.T) {
		um := newUnmarshaller(oneSpace)

		in := `key: value # comment`
		exp := `key: "value" # comment`

		out, err := um.unmarshallString(in)

		if err != nil {
			t.Errorf("Error: %s", err)
			return
		}

		testEqual(t, exp, out)
	})

	t.Run("Simple boolean key value", func(t *testing.T) {
		um := newUnmarshaller(oneSpace)

		in := `key: true`
		exp := `key: true`
		out, err := um.unmarshallString(in)

		if err != nil {
			t.Errorf("Error: %s", err)
			return
		}

		testEqual(t, exp, out)
	})

	t.Run("Simple boolean key value with comment", func(t *testing.T) {
		um := newUnmarshaller(oneSpace)

		in := `key: true # comment`
		exp := `key: true # comment`
		out, err := um.unmarshallString(in)

		if err != nil {
			t.Errorf("Error: %s", err)
			return
		}

		testEqual(t, exp, out)
	})

	t.Run("Simple integer key value", func(t *testing.T) {
		um := newUnmarshaller(oneSpace)

		in := `key: 1`
		exp := `key: 1`
		out, err := um.unmarshallString(in)

		if err != nil {
			t.Errorf("Error: %s", err)
			return
		}

		testEqual(t, exp, out)
	})

	t.Run("Simple integer key value with comment", func(t *testing.T) {
		um := newUnmarshaller(oneSpace)

		in := `key: 1 # comment`
		exp := `key: 1 # comment`
		out, err := um.unmarshallString(in)

		if err != nil {
			t.Errorf("Error: %s", err)
			return
		}

		testEqual(t, exp, out)
	})

	t.Run("Simple float key value", func(t *testing.T) {
		um := newUnmarshaller(oneSpace)

		in := `key: 1.0`
		exp := `key: 1.0`
		out, err := um.unmarshallString(in)

		if err != nil {
			t.Errorf("Error: %s", err)
			return
		}

		testEqual(t, exp, out)
	})

	t.Run("Simple float key value with comment", func(t *testing.T) {
		um := newUnmarshaller(oneSpace)

		in := `key: 1.0 # comment`
		exp := `key: 1.0 # comment`
		out, err := um.unmarshallString(in)

		if err != nil {
			t.Errorf("Error: %s", err)
			return
		}

		testEqual(t, exp, out)
	})

	t.Run("Simple null key value", func(t *testing.T) {
		um := newUnmarshaller(oneSpace)

		in := `key: null`
		exp := `key: null`
		out, err := um.unmarshallString(in)

		if err != nil {
			t.Errorf("Error: %s", err)
			return
		}

		testEqual(t, exp, out)
	})

	t.Run("Simple null key value with comment", func(t *testing.T) {
		um := newUnmarshaller(oneSpace)

		in := `key: null # comment`
		exp := `key: null # comment`
		out, err := um.unmarshallString(in)

		if err != nil {
			t.Errorf("Error: %s", err)
			return
		}

		testEqual(t, exp, out)
	})

}

func TestUnmarshallerArray(t *testing.T) {

	t.Run("Simple array", func(t *testing.T) {
		um := newUnmarshaller(oneSpace)

		in := `
key:
  - 1
  - string
  - true
  - 3.1415
  - null`

		exp := `
key: [
 1,
 "string",
 true,
 3.1415,
 null
]`
		out, err := um.unmarshallString(in)

		if err != nil {
			t.Errorf("Error: %s", err)
			return
		}

		testEqual(t, exp, out)
	})

	t.Run("Simple array with comment", func(t *testing.T) {
		um := newUnmarshaller(oneSpace)

		in := `
# comment
key: # comment
  - 1 # comment
  - string # comment
  - true # comment
  - 3.1415 # comment
  - null # comment`

		exp := `
# comment
key: [ # comment
 1, # comment
 "string", # comment
 true, # comment
 3.1415, # comment
 null # comment
]`
		out, err := um.unmarshallString(in)

		if err != nil {
			t.Errorf("Error: %s", err)
			return
		}

		testEqual(t, exp, out)
	})

	t.Run("Array with simple object", func(t *testing.T) {

		um := newUnmarshaller(oneSpace)

		in := `
key:
  - obj:
       key: value`

		exp := `
key: [
 obj:
  {
   key: "value"
  }
]`
		out, err := um.unmarshallString(in)

		if err != nil {
			t.Errorf("Error: %s", err)
			return
		}

		testEqual(t, exp, out)

	})
}

func TestUnmarshallerObject(t *testing.T) {

	t.Run("Simple object", func(t *testing.T) {
		um := newUnmarshaller(oneSpace)

		in := `
key:
  key: value`

		exp := `
key: {
 key: "value"
}`
		out, err := um.unmarshallString(in)

		if err != nil {
			t.Errorf("Error: %s", err)
			return
		}

		testEqual(t, exp, out)
	})

	t.Run("Simple object with comment", func(t *testing.T) {
		um := newUnmarshaller(oneSpace)

		in := `
# comment
key: # comment
  key: value # comment`

		exp := `
# comment
key: { # comment
 key: "value" # comment
}`
		out, err := um.unmarshallString(in)

		if err != nil {
			t.Errorf("Error: %s", err)
			return
		}

		testEqual(t, exp, out)
	})

	t.Run("Object with simple array", func(t *testing.T) {
		um := newUnmarshaller(oneSpace)

		in := `
key:
  key:
	- value`

		exp := `
key: {
 key: [
  "value"
 ]
}`
		out, err := um.unmarshallString(in)

		if err != nil {
			t.Errorf("Error: %s", err)
			return
		}

		testEqual(t, exp, out)
	})
}

func TestUnmarshallerLiteral(t *testing.T) {

	t.Run("Simple literal", func(t *testing.T) {
		um := newUnmarshaller(oneSpace)

		in := `
literal : |
 this is literal
 right?
`
		exp := `
literal: |
 this is literal
 right?
`
		out, err := um.unmarshallString(in)

		if err != nil {
			t.Errorf("Error: %s", err)
			return
		}

		testEqual(t, exp, out)
	})
}
