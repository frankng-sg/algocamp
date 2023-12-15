package main

import (
	"strconv"
	"strings"
)

func decode(line string) (record []RecordType, reqs []int) {
	parts := strings.Split(line, " ")
	record = []RecordType(parts[0])
	strNums := strings.Split(parts[1], ",")
	for _, strNum := range strNums {
		num, _ := strconv.Atoi(strNum)
		reqs = append(reqs, num)
	}
	return record, reqs
}

func isEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func isPossibleCombo(record []RecordType, reqs []int) bool {
	var score []int
	count := 0
	for i := range record {
		if record[i] == DAMAGED {
			count++
		} else if count > 0 {
			score = append(score, count)
			count = 0
		}
	}
	if count > 0 {
		score = append(score, count)
	}
	return isEqual(score, reqs)
}

func expandRecord(record []RecordType, sep RecordType, n int) []RecordType {
	var result []RecordType
	result = append(result, record...)
	for i := 1; i < n; i++ {
		result = append(result, sep)
		result = append(result, record...)
	}
	return result
}

func expandReqs(reqs []int, n int) []int {
	var result []int
	for i := 0; i < n; i++ {
		result = append(result, reqs...)
	}
	return result
}

func canBePlaced(record []RecordType, from, to int) bool {
	if from > 0 && record[from-1] == DAMAGED {
		return false
	}
	if to >= len(record) {
		return false
	}
	if to < len(record)-1 && record[to+1] == DAMAGED {
		return false
	}
	for i := from; i <= to; i++ {
		if record[i] == OPERATIONAL {
			return false
		}
	}
	return true
}

func noMoreDamaged(record []RecordType, from int) bool {
	for i := from; i < len(record); i++ {
		if record[i] == DAMAGED {
			return false
		}
	}
	return true
}
