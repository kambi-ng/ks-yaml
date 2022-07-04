package ksyaml

import "testing"
import "strings"

const oneSpace = " "

type testTable struct {
	name     string
	input    string
	expected string
}

func assertEqual(t *testing.T, expected, output string) {
	expected = strings.TrimSpace(expected)
	output = strings.TrimSpace(output)
	if expected != output {
		t.Errorf("\nExpected:\n\n%s\n\nGot:\n\n%s\n", expected, output)
	}
}

func TestUnmarshallerSimple(t *testing.T) {

	tb := []testTable{
		{
			name:     "Simple String Key Value",
			input:    `key: value`,
			expected: `key: "value"`,
		},
		{
			name:     "Simple String Key Value with comment",
			input:    `key: value # comment`,
			expected: `key: "value" # comment`,
		},

		{
			name:     "Simple boolean Key Value",
			input:    `key: true`,
			expected: `key: true`,
		},
		{
			name:     "Simple boolean Key Value with comment",
			input:    `key: true # comment`,
			expected: `key: true # comment`,
		},
		{
			name:     "Simple integer Key Value",
			input:    `key: 1`,
			expected: `key: 1`,
		},
		{
			name:     "Simple integer Key Value with comment",
			input:    `key: 1 # comment`,
			expected: `key: 1 # comment`,
		},
		{
			name:     "Simple float Key Value",
			input:    `key: 3.1415`,
			expected: `key: 3.1415`,
		},
		{
			name:     "Simple float Key Value with comment",
			input:    `key: 3.1415 # comment`,
			expected: `key: 3.1415 # comment`,
		},
		{
			name:     "Simple null Key Value",
			input:    `key: null`,
			expected: `key: null`,
		},
	}

	for _, tt := range tb {
		t.Run(tt.name, func(t *testing.T) {
			um := newUnmarshaller(oneSpace)
			out, err := um.unmarshallString(tt.input)
			if err != nil {
				t.Errorf("Error: %s", err)
				return
			}
			assertEqual(t, tt.expected, out)
		})
	}
}

func TestUnmarshallerArray(t *testing.T) {

	tb := []testTable{
		{
			name: "Simple array",
			input: `
key:
  - 1
  - string
  - true
  - 3.1415
  - null`,
			expected: `
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
			input: `
key: # comment
  - 1 # comment
  - string # comment
  - true # comment
  - 3.1415 # comment
  - null # comment`,
			expected: `
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
			input: `
key:
  - obj:
       key: value`,
			expected: `
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
			out, err := um.unmarshallString(tt.input)
			if err != nil {
				t.Errorf("Error: %s", err)
				return
			}
			assertEqual(t, tt.expected, out)
		})
	}
}

func TestUnmarshallerObject(t *testing.T) {

	tb := []testTable{
		{
			name: "Simple object",
			input: `
key:
  key: value`,
			expected: `
key: {
 "key": "value"
}`,
		},
		{
			name: "Simple object with comment",
			input: `
key: # comment
  key: value # comment`,
			expected: `
key: { # comment
 "key": "value" # comment
}`,
		},
		{
			name: "Object with array",
			input: `
key:
  arr:
	- 1
	- string
`,
			expected: `
key: {
 arr: [
  1,
  "string"
 ]
}`,
		},
		{
			name: "Nested Object",
			input: `
key:
  key:
	key: value`,
			expected: `
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
			out, err := um.unmarshallString(tt.input)
			if err != nil {
				t.Errorf("Error: %s", err)
				return
			}
			assertEqual(t, tt.expected, out)
		})
	}
}

func TestUnmarshallerLiteral(t *testing.T) {

	tb := []testTable{
		{
			name: "Simple literal",
			input: `
litral: |
  this is literal
  right?`,
			expected: `
litral: |
  this is literal
  right?`,
		},
	}
	for _, tt := range tb {
		t.Run(tt.name, func(t *testing.T) {
			um := newUnmarshaller(oneSpace)
			out, err := um.unmarshallString(tt.input)
			if err != nil {
				t.Errorf("Error: %s", err)
				return
			}
			assertEqual(t, tt.expected, out)
		})
	}
}
