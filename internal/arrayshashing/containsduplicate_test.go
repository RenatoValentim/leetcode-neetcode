package arrayshashing_test

import (
  "testing"
  "github.com/RenatoValentim/leetcode-neetcode/internal/arrayshashing"
  "github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
  assert := assert.New(t)

  testCases := []struct {
    name string
    input []int
    expected bool
  } {
      {
        name: "Should return false when providing an empty slice",
        input: []int{},
        expected: false,
      },
  }

  for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
      output := arrayshashing.HasDuplicated(tc.input)
      assert.Equal(tc.expected, output)
    })
  }
}

