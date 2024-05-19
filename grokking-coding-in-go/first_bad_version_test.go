package main

func isBadVersion(s int) bool {
	return true
}

func firstBadVersion(n int) (int, int) {
	l, r := 1, n
	cnt := 0
	for l <= r {
		m := (l + r) / 2
		cnt++
		if isBadVersion(m) {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l, cnt
}
