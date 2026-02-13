package parser

import "testing"

func TestSubstitute(t *testing.T) {
	script := "Hello {{NAME}}, welcome to {{PLACE}}!"
	vars := map[string]string{
		"NAME":  "Justin",
		"PLACE": "Doozers",
	}

	expected := "Hello Justin, welcome to Doozers!"
	actual := Substitute(script, vars)

	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
