package _0981

import (
	"runtime/debug"
	"sort"
)

type pair struct {
	timestamp int
	value     string
}

type TimeMap struct {
	m map[string][]pair
}

func Constructor() TimeMap {
	return TimeMap{map[string][]pair{}}
}

func (m *TimeMap) Set(key string, value string, timestamp int) {
	m.m[key] = append(m.m[key], pair{timestamp, value})
}

func (m *TimeMap) Get(key string, timestamp int) string {
	pairs := m.m[key]
	i := sort.Search(len(pairs), func(i int) bool { return pairs[i].timestamp > timestamp })
	if i > 0 {
		return pairs[i-1].value
	}
	return ""
}

func init() { debug.SetGCPercent(-1) }

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
