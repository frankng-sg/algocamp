package main

func arrmax(nums []int) int {
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

func eatingTimeInHours(piles []int, speed int) int {
	sum := 0
	for _, v := range piles {
		sum += (v-1)/speed + 1
	}
	return sum
}

func minEatingSpeed(piles []int, h int) int {
	l, r := 1, arrmax(piles)
	for l < r {
		mid := l + (r-l)>>1
		if eatingTimeInHours(piles, mid) <= h {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
