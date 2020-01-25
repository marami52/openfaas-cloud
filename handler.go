package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	axcelURL := "http://api.ironix.com"
	owner := "marami520"
	var err error

	c := http.Client{
		Timeout: time.Second * 3,
	}

	request, _ := http.NewRequest(http.MethodGet, axcelURL+"/api/auth/deployfunction/"+owner, nil)

	response, err := c.Do(request)
	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))
}
