package service

import (
	"sort"
	"time"
)

func SearchLowTime(list *map[string]time.Duration) (string, time.Duration) {
	var temp []time.Duration

	for _, dur := range *list {
		if dur > 0 {
			temp = append(temp, dur)
		}
	}

	sort.Slice(temp, func(i, j int) bool {
		return temp[i] < temp[j]
	})

	low := temp[0]

	var url string
	for key, val := range *list {
		if val == low {
			url = key
			break
		}
	}

	return url, low
}
