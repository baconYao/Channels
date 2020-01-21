package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// 建立一個 Channel，作為每個 routine 的溝通橋樑
	// channel內傳遞的資料type為string (communicate over with string)
	c := make(chan string)

	for _, link := range links {
		// 利用 keyword "go" 來建立 go routine
		// 將 channel "c" 傳入 function
		go checkLink(link, c)
	}
}

func checkLink(link string, c chan string) {
	// http.Get is Blocking call
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
}
