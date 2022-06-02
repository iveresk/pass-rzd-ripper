package main

import (
	"log"
	"time"
)

func main() {
	channel_count := 4000000
	// proxies := ?PROXY COUNT
	for {
		for i := 0; i < channel_count; i++ {
			ch := make(chan string)
			go httpTicket(ch)
			log.Println(<-ch)
		}
		time.Sleep(1 * time.Minute)
	}
}
