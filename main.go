package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

/*
	go routines는 메인 함수를 기다려주지 않음.
	따라서 go routines와 메인 함수와의 커뮤니케이션이 필요함.
	이 때 사용되는 것이 Channel -> pipe로 결과를 메인에 전달함
*/

func main() {
	c := make(chan string)
	var results = make(map[string]string) // 빈 맵 초기화
	urls := []string{
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.airbnb.com",
		"https://www.instagram.com",
		"https://www.udemy.com",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}
	for _, url := range urls {
		results[url] = <-c
	}
	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		c <- "FAILED"
	}
	c <- "OK"
}
