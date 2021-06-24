package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	api := "http://localhost:8080"
	longUrl := "https://test.com"

	data := url.Values{
		"url": {longUrl},
	}

	resp, err := http.PostForm(api + "/short", data)

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println("Short url:", res["url"])
}
