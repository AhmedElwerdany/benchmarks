package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	bytes, err := os.ReadFile("../data.json")
	if err != nil {
		panic(err)
	}
	var data = []byte(string(bytes))
	var result map[string]interface{}
	performance1 := time.Now()
	err = json.Unmarshal(data, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Parsing 10,000 field %v", time.Since(performance1))
}
