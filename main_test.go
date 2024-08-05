package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.")

	exp := 19

	res, err := count(b)
	if err != nil {
		t.Error(err)
	}

	if res != exp {
		t.Errorf("Expected %d, got %d instead.", exp, res)
	}
}
