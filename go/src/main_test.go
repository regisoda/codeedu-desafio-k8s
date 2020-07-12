package main

import "testing"

func TestGreeting(t *testing.T) {
	str := "XYZ"
	resultExpected := "<b>" + str + "</b>"
	var result = greeting(str)
	if result != resultExpected {
		t.Errorf("Expected: %s, returns: %s", resultExpected, result)
	}
}
