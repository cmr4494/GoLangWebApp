package main

import "testing"

func TestHellow(t *testing.T) {
	result := Hello()
	expected := "hello world"

	if result != expected {
		t.Errorf("result %q expected %q", result, expected)
	}
}
