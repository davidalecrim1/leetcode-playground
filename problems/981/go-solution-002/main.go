package main

type Entry struct {
	timestamp int
	value     string
}

type TimeMap struct {
	data map[string][]Entry
}

func Constructor() TimeMap {
	return TimeMap{
		data: make(map[string][]Entry),
	}
}

func (tm *TimeMap) Set(key string, value string, timestamp int) {
	tm.data[key] = append(tm.data[key], Entry{
		timestamp: timestamp,
		value:     value,
	})
}

func (tm *TimeMap) Get(key string, timestamp int) string {
	entries := tm.data[key]
	if len(entries) == 0 {
		return ""
	}

	left, right := 0, len(entries)-1
	res := ""

	for left <= right {
		mid := (left + right) / 2

		if entries[mid].timestamp <= timestamp {
			res = entries[mid].value
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return res
}
