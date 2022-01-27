package main

import (
	"Project1/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
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
	router := chi.NewRouter()
	router.Post("/text", func(writer http.ResponseWriter, request *http.Request) {
		jdata := struct {
			Text string `json :"text"`
		}{}

		if parseErr := utils.ParseBody(request.Body, &jdata); parseErr != nil {
			utils.RespondError(writer, http.StatusBadRequest, parseErr, "failed to parse request body")
			return
		}

		textjson, err := json.Marshal(jdata)
		fmt.Println(textjson)
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
		utils.RespondJSON(writer, http.StatusOK, p[0:11])

	})
	http.ListenAndServe(":8080", router)

}
