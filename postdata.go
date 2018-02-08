package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//PostDataWithResponse converts the input to json and then posts as body
func PostDataWithResponse(url string, body interface{}) (bool, []byte) {
	isQuite := false
	bytes, _ := json.MarshalIndent(body, "", "  ")
	if !isQuite {
		fmt.Printf("POST: %s\n", url)
		fmt.Printf("Request posted:\n%s\n", string(bytes))
	}
	resp, err := http.Post(url, "application/json", strings.NewReader(string(bytes)))
	if err != nil {
		fmt.Printf("Error in response %v\n", err)
		return false, nil
	}
	responseString, _ := ioutil.ReadAll(resp.Body)
	if !isQuite {
		fmt.Printf("Status : %s :\n", resp.Status)
		fmt.Printf("Response : %s :\n", responseString)

	}
	if resp.StatusCode == 200 {
		return true, responseString
	}
	return false, nil
}
