package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

func getStatus(name, url string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(response.Body)

	fmt.Println(name, response.StatusCode)
}

func main() {
	byteValue, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println(err)
	}
	var result map[string]string
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		fmt.Println(err)
	}
	var wg sync.WaitGroup
	for k, v := range result {
		wg.Add(1)
		go func(name, url string) {
			defer wg.Done()
			getStatus(name, url)
		}(k, v)
		wg.Wait()
		fmt.Println("Done")
	}
}
