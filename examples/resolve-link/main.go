package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	api:= "http://localhost:8080"
	shortUrl := "http://localhost:8080/1"

	data := url.Values{
		"url": {shortUrl},
	}

	resp, err := http.PostForm(api + "/long", data)

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println("Long url:", res["url"])
}
