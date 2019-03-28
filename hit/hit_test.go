package hit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeKey(t *testing.T) {
	testCase := []struct {
		label    string
		int1     int
		int2     int
		limit    int
		str1     string
		str2     string
		expected string
	}{
		{
			label:    "test 1",
			int1:     3,
			int2:     5,
			str1:     "fizz",
			str2:     "buzz",
			limit:    15,
			expected: "int1=3&int2=5&limit=15&str1=fizz&str2=buzz",
		},
	}

	for _, tc := range testCase {
		generatedKeyName := makeKey(tc.int1, tc.int2, tc.limit, tc.str1, tc.str2)
		assert.Equal(t, tc.expected, generatedKeyName, tc.label)
	}
}

func TestAdd(t *testing.T) {

	h := New()

	int1 := 3
	int2 := 5
	limit := 20
	str1 := "fizz"
	str2 := "fuzz"

	expectedMostFrequentedRequest := makeKey(int1, int2, limit, str1, str2)

	for i := 0; i < 2; i++ {
		h.Add(int1, int2, limit, str1, str2)

	}
	h.Add(1, 2, 3, "puzzle", "buble")

	assert.Equal(t, expectedMostFrequentedRequest, h.GetMostFrequentRequest(), "")
}
