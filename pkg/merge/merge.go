package merge

func Merge(left []int, right []int) []int {
	result := make([]int, len(left) + len(right))
	lidx, ridx := 0, 0

	for idx := 0; idx < len(result); idx++ {
		// Has left run out?
		if lidx >= len(left) {
			result[idx] = right[ridx]
			ridx++
		} else if ridx >= len(right) {
			result[idx] = left[lidx]
			lidx++
		} else if left[lidx] < right[ridx] {
			result[idx] = left[lidx]
			lidx++
		} else {
			result[idx] = right[ridx]
			ridx++
		}
	}

	return result
}

// Sort performs a concurrent merge sort using goroutines
func Sort(data []int, result chan []int) {
	if len(data) == 1 {
		result <- data
		return
	}
	// Split into two sub-parts and run them concurrently
	left := make(chan []int)
	right := make(chan []int)

	mid := uint(len(data) / 2)

	ldata := data[:mid]
	rdata := data[mid:]

	// Fork
	go Sort(ldata, left)
	go Sort(rdata, right)

	// Join
	ldata = <-left
	rdata = <-right

	result <- Merge(ldata, rdata)
	return
}
