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
		{
			name:     "Should return an empty array when providing a slice with len < 2",
			input1:   []int{1},
			input2:   10,
			expected: []int{},
		},
		{
			name:     "Should return a slice with of the two numbers that sum up to the target",
			input1:   []int{2, 7, 11, 15},
			input2:   9,
			expected: []int{0, 1},
		},
		{
			name:     "Should return a slice with of the two numbers that sum up to the target",
			input1:   []int{3, 2, 4},
			input2:   6,
			expected: []int{1, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := arrayshashing.TwoSum(tc.input1, tc.input2)
			assert.Equal(tc.expected, output)
		})
	}
}
