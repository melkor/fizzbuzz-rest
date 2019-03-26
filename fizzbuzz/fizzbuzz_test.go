package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {

	testCase := []struct {
		label    string
		int1     int
		int2     int
		limit    int
		str1     string
		str2     string
		expected []string
	}{
		{
			label: "classic FizzBuzz",
			int1:  3,
			int2:  5,
			str1:  "fizz",
			str2:  "buzz",
			limit: 15,
			expected: []string{
				"1",
				"2",
				"fizz",
				"4",
				"buzz",
				"fizz",
				"7",
				"8",
				"fizz",
				"buzz",
				"11",
				"fizz",
				"13",
				"14",
				"fizzbuzz",
			},
		},
	}

	for _, tc := range testCase {
		result, error := Fizzbuzz(tc.int1, tc.int2, tc.limit, tc.str1, tc.str2)
		assert.NoError(t, error)
		assert.Equal(t, tc.expected, result, tc.label)
	}
}
