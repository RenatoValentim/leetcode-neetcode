package arrayshashing_test

import (
	"testing"

	"github.com/RenatoValentim/leetcode-neetcode/internal/arrayshashing"
	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	assert := assert.New(t)

	testCases := []struct {
		name     string
		input1   []int
		input2   int
		expected []int
	}{
		{
			name:     "Should return an empty array when providing an empty slice",
			input1:   []int{},
			input2:   10,
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := arrayshashing.TwoSum(tc.input1, tc.input2)
			assert.Equal(tc.expected, output)
		})
	}
}
