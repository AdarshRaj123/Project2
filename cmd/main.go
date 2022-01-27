package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type WordCount struct {
	Word  string
	Count int
}

type WordCountList []WordCount

func main() {
	text := struct {
		Text string `json:"text"`
	}{Text: "he is a good boy very good boy he loves music and loves to be done not learning to go out"}
	textjson, err := json.Marshal(text)
	req, err := http.NewRequest("POST", "http://localhost:3000/text", bytes.NewBuffer(textjson))
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	req.Header.Add("Accept", `application/json`)

	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("Error %s", err)
		return

	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	p := make(WordCountList, 10)
	err1 := json.Unmarshal(body, &p)
	if err1 != nil {

	}
	fmt.Println("Body : %s \n ", p)

}
