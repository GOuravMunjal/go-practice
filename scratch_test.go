package main

import (
	"testing"
)

func TestGetNamelyID(t *testing.T) {
	input := "4432d92c-e161-4555-bc8f-4e096f1e7312"
	expectedOutput := employee{
		id:     "4432d92c-e161-4555-bc8f-4e096f1e7312",
		name:   "Shikhar",
		salary: 5000,
	}
	output, err := getNamelyID(input)

	if output != expectedOutput || err != nil {
		t.Fatalf("Test case Failed dear")
	}

	input = "4432d92c-e161-4555-bc8f-4e096f1e7313" //wrong id wala case
	output, err = getNamelyID(input)
	if err == nil {
		t.Fatalf("Test case Failed dear")
	}

}
