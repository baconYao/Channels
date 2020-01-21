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

	// 建立一個 Channel，作為每個 routine 的溝通橋樑 (注意，此channel不只可以讓 child routine和main routine溝通，child routine之間也可以互相溝通)
	// channel內傳遞的資料type為string (communicate over with string)
	c := make(chan string)

	for _, link := range links {
		// 利用 keyword "go" 來建立 go routine
		// 將 channel "c" 傳入 function
		go checkLink(link, c)
	}

	// 從channel接收被傳遞的訊息
	fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	// http.Get is Blocking call
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// 傳遞訊息給channel (注意: fmt.Println(link, "might be down!") 也會一起傳遞)
		c <- "Might be down I think!"
		return
	}

	fmt.Println(link, "is up!")
	// 傳遞訊息給channel (注意: fmt.Println(link, "is up!") 也會一起傳遞)
	c <- "Yep it's up"
}
