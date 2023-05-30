package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func makeRequest(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}
	bodyString := string(bodyBytes)

	return bodyString
}

func main() {
	var performaceArray [100]int

	for i := 0; i < 100; i++ {
		timeStart := time.Now()
		url := "https://example.com"
		makeRequest(url)
		performaceArray[i] = int(time.Since(timeStart).Milliseconds())
	}

	var average int = 0
	for _, value := range performaceArray {
		fmt.Printf("%d\n", value)
		average += value
	}

	fmt.Printf("Average: %d", average/100)
}
