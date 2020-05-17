package requesties

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	response string
}

func Get(url string) *Response {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	bodystring := string(body)
	return &Response{response: bodystring}
}
