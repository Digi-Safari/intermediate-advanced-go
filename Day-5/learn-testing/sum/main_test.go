package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// add _test at the end of the fileName to make it a test file
// go test ./... // run all the test for the project

// function names must start with word Test to signal it is a test
// helper functions could be present in this file
// that would not be part of test so the would not start with the word Test
func TestSumInt(t *testing.T) {
	// Figure out two things
	// What are inputs (parameters)
	// What are outputs (return values)

	input := []int{1, 2, 3, 4, 5}
	want := 15
	got := SumInt(input)
	if got != want {
		// test would continue on if test case fail
		t.Errorf("sum of 1 to 5 should be %v; got %v", want, got)

		// Uncomment next line to stop the test if it fails at this point.
		//t.Fatalf("sum of 1 to 5 should be %v; got %v", want, got)

	}

	want = 0
	got = SumInt(nil)

	if got != want {
		t.Errorf("sum of nil should be %v; got %v", want, got)
	}

}

// Table test
func TestTableTestSumInt(t *testing.T) {
	//type args struct {
	//	a int
	//	s string
	//}
	tt := [...]struct {
		name  string // name of the test
		input []int
		want  int
	}{
		{
			name:  "one to five numbers",
			input: []int{1, 2, 3, 4, 5},
			want:  15,
		},
		{
			name:  "nil slice",
			input: nil,
			want:  0,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := SumInt(tc.input)
			require.Equal(t, tc.want, got)
			//if got != tc.want {
			//	//t.Fatalf("sum of %v should be %v; got %v", tc.input, tc.want, got)
			//
			//}
		})

	}

}
