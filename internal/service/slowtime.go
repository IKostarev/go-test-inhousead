package service

import (
	"sort"
	"time"
)

func SearchSlowTime(list *map[string]time.Duration) (string, time.Duration) {
	var temp []time.Duration

	for _, dur := range *list {
		temp = append(temp, dur)
	}

	sort.Slice(temp, func(i, j int) bool {
		return temp[i] < temp[j]
	})

	slow := temp[len(temp)-1]

	var url string
	for key, val := range *list {
		if val == slow {
			url = key
			break
		}
	}

	return url, slow
}
