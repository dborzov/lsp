package main

import "testing"

func TestParseArguments(t *testing.T) {
	in := []string{"lsp", "-a"}

	x, err := ParseArguments(in)

	if err != nil {
		t.Errorf("ParseAguments(%#v) = %#v, %#v ", in, x, err)
		return
	}

}
