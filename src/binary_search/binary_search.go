package binary_search

func Search(numbers []int, number int) int {
  for low, high := 0, len(numbers)-1 ; low <= high; {
    mid := (low + high) / 2
    if number == numbers[mid] {
      return mid
    } else if number < numbers[mid] {
      high = mid - 1
    } else {
      low = mid + 1
    }
  }
  return -1
}
