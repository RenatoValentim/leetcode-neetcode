package arrayshashing_test

import (
  "testing"
  "github.com/RenatoValentim/leetcode-neetcode/internal/arrayshashing"
  "github.com/stretchr/testify/assert"
)

func TestValidAnagram(t *testing.T) {
  assert := assert.New(t)

  testCases := []struct {
    name string
    input1 string
    input2 string
    expected bool
  } {
      {
        name: "Should return false when providing any empty string",
        input1: "",
        input2: "",
        expected: false,
      },
      {
        name: "Should return false when length is not equals",
        input1: "a",
        input2: "aa",
        expected: false,
      },
  }

  for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
      output := arrayshashing.IsAnagram(tc.input1, tc.input2)
      assert.Equal(tc.expected, output)
    })
  }
}
