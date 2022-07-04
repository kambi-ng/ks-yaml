package ksyaml

import "testing"
import "strings"

const oneSpace = " "

type testTable struct {
	name string
	in   string
	exp  string
}

func assertEqual(t *testing.T, exp, out string) {
	exp = strings.TrimSpace(exp)
	out = strings.TrimSpace(out)
	if exp != out {
		t.Errorf("Expected:\n%s\n\nGot:\n\n%s", exp, out)
	}
}

func TestUnmarshallerSimple(t *testing.T) {

	tb := []testTable{
		{
			name: "Simple String Key Value",
			in:   `key: value`,
			exp:  `key: "value"`,
		},
		{
			name: "Simple String Key Value with comment",
			in:   `key: value # comment`,
			exp:  `key: "value" # comment`,
		},

		{
			name: "Simple boolean Key Value",
			in:   `key: true`,
			exp:  `key: true`,
		},
		{
			name: "Simple boolean Key Value with comment",
			in:   `key: true # comment`,
			exp:  `key: true # comment`,
		},
		{
			name: "Simple integer Key Value",
			in:   `key: 1`,
			exp:  `key: 1`,
		},
		{
			name: "Simple integer Key Value with comment",
			in:   `key: 1 # comment`,
			exp:  `key: 1 # comment`,
		},
		{
			name: "Simple float Key Value",
			in:   `key: 3.1415`,
			exp:  `key: 3.1415`,
		},
		{
			name: "Simple float Key Value with comment",
			in:   `key: 3.1415 # comment`,
			exp:  `key: 3.1415 # comment`,
		},
		{
			name: "Simple null Key Value",
			in:   `key: null`,
			exp:  `key: null`,
		},
	}

	for _, tt := range tb {
		t.Run(tt.name, func(t *testing.T) {
			um := newUnmarshaller(oneSpace)
			out, err := um.unmarshallString(tt.in)
			if err != nil {
				t.Errorf("Error: %s", err)
				return
			}
			assertEqual(t, tt.exp, out)
		})
	}
}

func TestUnmarshallerArray(t *testing.T) {

	tb := []testTable{
		{
			name: "Simple array",
			in: `
key:
  - 1
  - string
  - true
  - 3.1415
  - null`,
			exp: `
key: [
 1,
 string,
 true,
 3.1415,
 null
]`,
		},
		{
			name: "Simple array with comment",
			in: `
key: # comment
  - 1 # comment
  - string # comment
  - true # comment
  - 3.1415 # comment
  - null # comment`,
			exp: `
key: [ # comment
 1, # comment
 "string", # comment
 true, # comment
 3.1415, # comment
 null # comment
]`,
		},
		{
			name: "Array with object",
			in: `
key:
  - obj:
       key: value`,
			exp: `
key: [
 {
   "key": "value"
 }
]`,
		},
	}

	for _, tt := range tb {
		t.Run(tt.name, func(t *testing.T) {
			um := newUnmarshaller(oneSpace)
			out, err := um.unmarshallString(tt.in)
			if err != nil {
				t.Errorf("Error: %s", err)
				return
			}
			assertEqual(t, tt.exp, out)
		})
	}
}

func TestUnmarshallerObject(t *testing.T) {

	tb := []testTable{
		{
			name: "Simple object",
			in: `
key:
  key: value`,
			exp: `
key: {
 "key": "value"
}`,
		},
		{
			name: "Simple object with comment",
			in: `
key: # comment
  key: value # comment`,
			exp: `
key: { # comment
 "key": "value" # comment
}`,
		},
		{
			name: "Object with array",
			in: `
key:
  arr:
	- 1
	- string
`,
			exp: `
key: {
 arr: [
  1,
  "string"
 ]
}`,
		},
		{
			name: "Nested Object",
			in: `
key:
  key:
	key: value`,
			exp: `
key: {
 "key": {
  "key": "value"
 }
}`,
		},
	}

	for _, tt := range tb {
		t.Run(tt.name, func(t *testing.T) {
			um := newUnmarshaller(oneSpace)
			out, err := um.unmarshallString(tt.in)
			if err != nil {
				t.Errorf("Error: %s", err)
				return
			}
			assertEqual(t, tt.exp, out)
		})
	}
}

func TestUnmarshallerLiteral(t *testing.T) {

	tb := []testTable{
		{
			name: "Simple literal",
			in: `
litral: |
  this is literal
  right?`,
			exp: `
litral: |
  this is literal
  right?`,
		},
	}
	for _, tt := range tb {
		t.Run(tt.name, func(t *testing.T) {
			um := newUnmarshaller(oneSpace)
			out, err := um.unmarshallString(tt.in)
			if err != nil {
				t.Errorf("Error: %s", err)
				return
			}
			assertEqual(t, tt.exp, out)
		})
	}
}
