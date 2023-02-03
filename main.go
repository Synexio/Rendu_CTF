package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
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
	//fmt.Println(url)
	client := http.Client{
		Timeout: time.Second * 5,
	}
	resp, err := client.Get(url)
	//resp, err := http.Get(url)
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

func secondAPI(secret string, c2 chan string) {
	url := fmt.Sprintf("http://%s:%d", ip, secondPort)
	fmt.Println(url)
	url2 := fmt.Sprintf("http://%s:%d?secretKey=%s", ip, secondPort, secret)
	fmt.Println(url2)

	/*postBody, _ := json.Marshal(map[string]string{
		"secretKey": secret,
	})
	responseBody := bytes.NewBuffer(postBody)*/

	responseBody := []byte(`{
		"secretKey": "ca32652906af8dd747e741cd3e960338138099b0615e62b4f23366cf65f52646"
	}`)

	/*param := url.Values{}
	param.Add("secretKey", secret)
	data := url.Values{
	        "secretKey": {secret},
	    }
	http.PostForm(url,param)*/

	/*values := map[string]string{"secretKey": secret}
	json_data, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post(url, "application/json",
		bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	fmt.Println(res["json"])*/

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(responseBody))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	var wg sync.WaitGroup
	for port := 3000; port <= 4000; port++ {
		wg.Add(1)
		go firstAPI(port, &wg, c1)
	}

	string := <-c1
	key := string[19:]
	fmt.Println(key)
	secondAPI(key, c2)

	string2 := <-c2
	key2 := string2[20:]
	fmt.Println(key2)
}