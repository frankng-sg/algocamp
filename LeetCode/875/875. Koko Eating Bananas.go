package main

func arrsum(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func arrmax(nums []int) int {
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

func eatingTime(piles []int, speed int) int {
	sum := 0
	for _, v := range piles {
		sum += (v-1)/speed + 1
	}
	return sum
}

func minEatingSpeed(piles []int, h int) int {
	var maxspeed, minspeed int
	if sum := arrsum(piles); sum%h == 0 {
		minspeed = sum / h
	} else {
		minspeed = sum/h + 1
	}
	maxspeed = arrmax(piles)
	l, r := minspeed, maxspeed
	for l < r {
		mid := l + (r-l)>>1
		if eatingTime(piles, mid) <= h {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
