package ksyaml

import "testing"
import "strings"

func equalTrimmed(a, b string) bool {
	return strings.TrimSpace(a) == strings.TrimSpace(b)
}

func testEqual(t *testing.T, exp, out string) {
	if !equalTrimmed(exp, out) {
		t.Errorf("Expected: \n %s \n Got: \n %s", exp, out)
	}
}

func TestUnmarshallerSimple(t *testing.T) {
	um := newUnmarshaller(" ")
	t.Run("Simple string key value", func(t *testing.T) {

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
