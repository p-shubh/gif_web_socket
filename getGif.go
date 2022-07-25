package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getGIF(query string) string {

	url := "https://g.tenor.com/v1/search?q=" + query + "&key=LIVDSRZULELA&limit=1"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return "unable to generate image"
	}
	req.Header.Add("apiKey", "0UTRbFtkMxAplrohufYco5IY74U8hOes")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "unable to generate image"
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "unable to generate image"
	}
	data := GIF{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return "unable to generate image"
	}
	return data.Results[0].Media[0].Gif.URL
}
