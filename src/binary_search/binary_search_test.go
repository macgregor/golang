package binary_search

import "testing"

func TestHappyPath(t *testing.T) {
  numbers := []int{1, 2, 3, 4, 5}
  want := 2
  got := Search(numbers, 3)

  if want != got {
    t.Errorf("Got %d, want %d", got, want)
  }
}

func TestLeftEdge(t *testing.T) {
  numbers := []int{1, 2, 3, 4, 5}
  want := 0
  got := Search(numbers, 1)

  if want != got {
    t.Errorf("Got %d, want %d", got, want)
  }
}

func TestRightEdge(t *testing.T) {
  numbers := []int{1, 2, 3, 4, 5}
  want := 4
  got := Search(numbers, 5)

  if want != got {
    t.Errorf("Got %d, want %d", got, want)
  }
}

func TestNegative(t *testing.T) {
  numbers := []int{1, 2, 3, 4, 5}
  want := -1
  got := Search(numbers, 0)

  if want != got {
    t.Errorf("Got %d, want %d", got, want)
  }
}

func TestAll(t *testing.T) {
  numbers := []int{1, 2, 3, 4, 5, 6}

  for i := 0; i < len(numbers); i = i + 1 {
    want := i
    got := Search(numbers, numbers[i])
    if want != got {
      t.Errorf("Got %d, want %d", got, want)
    }
  }
}
