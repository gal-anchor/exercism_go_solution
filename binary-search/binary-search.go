package binarysearch

// SearchInts searches for key in a sorted list of integers.
// It returns the index of key if found, otherwise -1.
func SearchInts(list []int, key int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := low + (high-low)/2 // avoids potential overflow

		if list[mid] == key {
			return mid
		} else if list[mid] < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1 // not found
}
