package service

import (
	"fmt"
	"net/http"
	"time"
)

func Request(lists *map[string]time.Duration) {
	ch := make(chan string)

	for url := range *lists {
		go makeRequest(lists, url, ch)
	}

	for range *lists {
		result := <-ch
		if result != "" {
			fmt.Println(result)
		}
	}
}

func makeRequest(lists *map[string]time.Duration, url string, ch chan<- string) {
	client := &http.Client{}
	start := time.Now()
	req, err := http.NewRequest("GET", "http://"+url, nil)
	if err != nil {
		ch <- ""
		return
	}

	resp, err := client.Do(req)
	duration := time.Since(start)
	if err == nil && (duration > time.Microsecond*5) && (duration < time.Second*5) {
		ch <- fmt.Sprintf("метод - %s | статус ответа - %d | к сайту - %s | длительность обращения - %v", req.Method, resp.StatusCode, req.URL, duration)
		(*lists)[url] = duration
	} else {
		ch <- ""
	}
}
