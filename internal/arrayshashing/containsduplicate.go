package arrayshashing

type Set map[int]bool

func (s Set) add(value int) {
  s[value] = true
}

func HasDuplicated(values []int) bool {
  if len(values) == 0 {
    return false
  }
  s := make(Set)
  for _, val := range values {
    s.add(val)
  }
  if len(values) != len(s) {
    return true
  }
  return false
}
