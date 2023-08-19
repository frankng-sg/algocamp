package main

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */

type TimeBasedValue struct {
	Value     string
	Timestamp int
}

type TimeMap struct {
	cache map[string][]TimeBasedValue // cache[key] = [[value, timestamp], ...]
}

func Constructor() TimeMap {
	return TimeMap{cache: make(map[string][]TimeBasedValue)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.cache[key] = append(this.cache[key], TimeBasedValue{value, timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	i := bsearch(this.cache[key], timestamp)
	if i == -1 {
		return ""
	}
	return this.cache[key][i].Value
}

func bsearch(data []TimeBasedValue, target int) int {
	l, r := 0, len(data)-1
	if r < 0 {
		return -1
	}
	for l < r-1 {
		mid := (l + r) >> 1
		if data[mid].Timestamp >= target {
			r = mid
		} else {
			l = mid
		}
	}
	if data[r].Timestamp <= target {
		return r
	}
	if data[l].Timestamp <= target {
		return l
	}
	return -1
}
