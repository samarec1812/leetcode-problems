package first_bad_version

func isBadVersion(n int) bool {
	// some actions
	return true
}

func firstBadVersion(n int) int {
	left, right := 0, n-1

	for left <= right {
		middle := (left + right) / 2
		if !isBadVersion(middle) {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return left
}
