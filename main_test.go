package main

import "testing"

// TODO: Please write real tests. Including integration tests regarding mongo.
// Mock mongo.

func TestHelloWorld(t *testing.T) {
	str := "eldiario"
	if str != "eldiario" {
		t.Errorf("Need more tests.")
	}
}
