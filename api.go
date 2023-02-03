package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	ip         = "34.77.36.161"
	secondPort = 3941
)

func firstAPI(port int, wg *sync.WaitGroup, c1 chan string) {
	defer wg.Done()
	url := fmt.Sprintf("http://%s:%d", ip, port)
	client := http.Client{
		Timeout: time.Second * 5,
	}
	resp, err := client.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		body, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			return
		}
		fmt.Printf("Response from port %d: %s\n", port, resp.Status)
		c1 <- string(body)
	}
}

func secondAPI(secret string) string {
	url := fmt.Sprintf("http://%s:%d", ip, secondPort)

	responseBody := []byte(fmt.Sprintf("secretKey=%s", secret))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(responseBody))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)

}

func thirdAPI() {
	/*content, err := ioutil.ReadFile("finalResult.txt")
	if err != nil {
		log.Fatal(err)
	}
	res := strings.Split(string(content), "\\n")
	thirdPort := res[0]
	key := res[1]
	value := res[2]*/

	//url := fmt.Sprintf("http://%s:%s", ip, thirdPort)
	url := fmt.Sprintf("http://%s:%s", ip, "3610")
	//responseBody := []byte(fmt.Sprintf("%s=%s", key, value))
	responseBody := []byte("finalKey=8116fdd3f12b6d7c4b136cbdaa3360a57eb4eb676ae63294450ee1f4f34b36f3")

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(responseBody))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client2 := &http.Client{}
	resp, err := client2.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("The flag is :", strings.Split(string(body), ":")[1][1:])
}