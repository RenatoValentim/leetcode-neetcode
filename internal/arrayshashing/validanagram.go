package arrayshashing

import (
  "strings"
  "sort"
)

func IsAnagram(input1, input2 string) bool {
  word1 := strings.ToLower(input1)
  word2 := strings.ToLower(input2)
  invalidInputs := word1 == "" || word2 == "" || len(word1) != len(word2) || len(word1) <= 1 && len(word2) >= 50000
  if invalidInputs {
    return false
  }
  w1 := strings.Split(word1, "")
  w2 := strings.Split(word2, "")
  sort.Strings(w1)
  sort.Strings(w2)
  return strings.Join(w1, "") == strings.Join(w2, "")
}
