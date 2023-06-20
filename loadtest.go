package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	url        = "http://127.0.0.1:8080/log"
	reqsPerSec = 1000
)

func generateRandomData() string {
	id := rand.Intn(10000)
	unixTs := 1609459200 + rand.Intn(1000000)
	userID := 100000 + rand.Intn(10000)
	eventName := "login"
	switch rand.Intn(4) {
	case 1:
		eventName = "logout"
	case 2:
		eventName = "view"
	case 3:
		eventName = "click"
	}

	data := fmt.Sprintf(`{
		"id": %d,
		"unix_ts": %d,
		"user_id": %d,
		"event_name": "%s"
	}`, id, unixTs, userID, eventName)

	return data
}

func sendRequest(data string, wg *sync.WaitGroup) {
	defer wg.Done()

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	_, err = client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		sleepDuration := time.Second / time.Duration(reqsPerSec)
		startTime := time.Now()
		var wg sync.WaitGroup

		for i := 0; i < reqsPerSec; i++ {
			wg.Add(1)
			go sendRequest(generateRandomData(), &wg)
			time.Sleep(sleepDuration - time.Since(startTime))
		}

		wg.Wait()
	}
}
