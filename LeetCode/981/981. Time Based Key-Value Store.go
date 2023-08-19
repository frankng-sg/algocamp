package main

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */

type TimeMap struct {
	values    map[string][]string
	timestamp map[string][]int
}

func Constructor() TimeMap {
	return TimeMap{
		values:    make(map[string][]string),
		timestamp: make(map[string][]int),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.values[key] = append(this.values[key], value)
	this.timestamp[key] = append(this.timestamp[key], timestamp)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	i := bsearch(this.timestamp[key], timestamp)
	if i == -1 {
		return ""
	}
	return this.values[key][i]
}

func bsearch(data []int, target int) int {
	l, r := 0, len(data)-1
	if r < 0 {
		return -1
	}
	for l < r-1 {
		mid := (l + r) >> 1
		if data[mid] >= target {
			r = mid
		} else {
			l = mid
		}
	}
	if data[r] <= target {
		return r
	}
	if data[l] <= target {
		return l
	}
	return -1
}
