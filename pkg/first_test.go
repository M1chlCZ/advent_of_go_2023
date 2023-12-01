package pkg

import (
	"testing"
)

func TestFindNumberPairs(t *testing.T) {
	t.Logf("\n%s", "Advent #1: ")
	testCases := []struct {
		input  string
		output int
	}{
		{
			input: `
	1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet`,
			output: 142,
		},
	}

	t.Logf("\nInput:%s, \nExpected: %d", testCases[0].input, testCases[0].output)

	for _, tc := range testCases {
		result, err := findNumberPairs(tc.input)
		if err != nil {
			t.Errorf("Expected no error but got %v", err)
		}

		sum := 0
		for _, num := range result {
			sum += num
		}
		t.Logf("\nResult: %d", sum)

		if sum != tc.output {
			t.Errorf("Expected sum %v but got %v", tc.output, sum)
		}
	}
}
