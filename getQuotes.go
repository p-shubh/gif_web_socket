package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getQuote() string {

	url := "https://zenquotes.io/api/random"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return "unable to generate quote"
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "unable to generate quote"
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "unable to generate quote"
	}
	data := []Quote{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return "unable to generate quote"
	}
	return data[0].Q
}
