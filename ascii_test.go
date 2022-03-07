package tools

import "testing"

var inputs = [3]string{
	"Bär",
	"Français",
	"être ou ne pas être",
}

func TestToAscii(t *testing.T) {
	outputs := [3]string{
		"Bar",
		"Francais",
		"etre ou ne pas etre",
	}
	for idx, input := range inputs {
		out := ToAscii(input)
		if out != outputs[idx] {
			t.Errorf("%q is not %q converted to pure ASCII", out, input)
		} else {
			t.Logf("%q is correctly converted to pure ASCII: %q", out, input)
		}
	}
}

func TestToAsciiSmall(t *testing.T) {
	outputs := [3]string{
		"bar",
		"francais",
		"etre ou ne pas etre",
	}
	for idx, input := range inputs {
		out := ToAsciiSmall(input)
		if out != outputs[idx] {
			t.Errorf("%q is not %q converted to pure ASCII Small", out, input)
		} else {
			t.Logf("%q is correctly converted to pure ASCII Small: %q", out, input)
		}
	}
}
