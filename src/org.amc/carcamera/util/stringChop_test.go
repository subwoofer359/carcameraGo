package util

import (
	"testing"
)

var testStr string = "Hello World How are you"

func TestStringLessThanN(t *testing.T) {
	strings := StringChop(testStr, 24)
	if strings[0] != testStr {
		t.Error("The function should return the same string")
	}
}

func TestStringGreaterThan2(t *testing.T) {
	size := 12
	strings := StringChop(testStr, size)
	if len(strings) != 2 {
		t.Fatal("The string array should contain two strings")
	}
	if strings[0] != testStr[0:size] {
		t.Fatal("The function should return the same string")
	}

	if strings[1] != testStr[size:] {
		t.Fatal("The function should return the same string")
	}
}

func TestStringGreaterThanN(t *testing.T) {
	size := 4
	noOfStrings := 6
	strings := StringChop(testStr, size)
	if len(strings) != noOfStrings {
		t.Fatalf("The string array should contain %d strings", noOfStrings)
	}

}

func TestStringEmpty(t *testing.T) {
	strings := StringChop("", 24)
	if len(strings) != 1 {
		t.Fatal("String array should have size of one")
	}

	if strings[0] != "" {
		t.Fatal("The null string should be in the array")
	}
}
