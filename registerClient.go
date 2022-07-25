package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsupgraders = websocket.Upgrader{
	ReadBufferSize:  1024 * 2,
	WriteBufferSize: 1024 * 2,
}

func RegisterClient(c *gin.Context) {

	wsupgraders.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := wsupgraders.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		msg := fmt.Sprintf("Failed to set websocket upgrade: %+v", err)
		fmt.Println(msg)
		return
	}

	for i := 0; i < 10; i++ {
		mType, mByte, err := conn.ReadMessage()
		fmt.Println("mByte: ", string(mByte))
		fmt.Println("mType: ", mType)
		fmt.Println("err: ", err)

		if string(mByte) != "quote" {
			link := getGIF(string(mByte))
			fmt.Println("link: ", link)
			image, err := downloadImage(link)
			if err == nil {
				conn.WriteMessage(websocket.BinaryMessage, image)
			} else {
				fmt.Println("error: ", err)
				conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%s", "unable to load image")))
			}
		} else {
			conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%s", getQuote())))
		}

	}
	conn.Close()
}

//download image from url

func downloadImage(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
