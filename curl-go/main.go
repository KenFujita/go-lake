package main

import (
	"fmt"
	"log"
	"net/http"
	"io"
)

func main(){
	url := "https://httpbin.org/get"

	// GETリクエストを送り、レスポンスを受け取る
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("not get response: %s\n",err)
	}

	//defer resp.Body.Close()

	// 受け取ったリクエストからBody部を取る
	body,err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("not get response: %s\n",err)
	}
	resp.Body.Close()

	fmt.Printf("%s\n",body)
}
