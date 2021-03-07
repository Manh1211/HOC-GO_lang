package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close() //make sure the body close

	ctype := resp.Header.Get("Conten-Type")
	if ctype == "" {
		//retrun error if content-type header not found
		return "", fmt.Errorf("can't find content_type header")
	}
	return ctype, nil
}

func main(){
	ctype, err := contentType("https://golang.org")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}else{
		fmt.Println(ctype)
	}
}