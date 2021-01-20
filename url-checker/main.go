package urlchecker

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

/*
	go routines는 메인 함수를 기다려주지 않음.
	따라서 go routines와 메인 함수와의 커뮤니케이션이 필요함.
	이 때 사용되는 것이 Channel -> pipe로 결과를 메인에 전달함
*/

func main() {
	c := make(chan requestResult)
	var results = make(map[string]string) // 빈 맵 초기화
	urls := []string{
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.airbnb.com",
		"https://www.dankook.ac.kr",
		"https://www.zini.run",
		"https://zinirun.github.io",
		"https://www.kakaocorp.com",
		"https://www.naver.com",
		"https://www.woowabros.com",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}
	for range urls {
		result := <-c
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
}
