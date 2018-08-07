package main

import (
	"testing"
	"fmt"
)

// go test -v (verbose)
// go test --cover >>
func TestPerson_FullName(t *testing.T) {

	tests := []struct{
		Input Person
		Output string
	}{
		{
			Input: Person{
				Name: "R",
				Surname: "L",
			},
			Output: "R L",
		},
		{
			Input: Person{
				Name: "M",
				Surname: "A",
			},
			Output: "M A",
		},
		{
			Input: Person{
				Name: "M",
				Surname: "",
			},
			Output: "M ",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			if test.Input.FullName() != test.Output {
				t.Errorf("Expected %s but got %s", "R L", test.Input.FullName())
			}
		})
	}
}