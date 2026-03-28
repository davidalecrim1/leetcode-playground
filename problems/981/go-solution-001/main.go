package main

// hashmap for the stored values
// the hashmap value would be a slice ordered by the timestamp

type Value struct {
	value     string
	timestamp int
}

type TimeMap struct {
	hm map[string][]Value
}

func Constructor() TimeMap {
	return TimeMap{
		hm: make(map[string][]Value, 0),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	val := Value{
		value:     value,
		timestamp: timestamp,
	}

	this.hm[key] = append(this.hm[key], val)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	values := this.hm[key]

	for i := len(values) - 1; i >= 0; i-- {
		val := values[i]
		if timestamp >= val.timestamp {
			return val.value
		}
	}

	return ""
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
