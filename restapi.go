package namegenerator

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// request makes a (GET) Requests to namegenerator2 REST API.
func request(urlStr string) (response []byte, err error) {
	c := http.Client{}

	resp, err := c.Get(urlStr)

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}

// parseRequest willload the request into a slice of strings
func parseRequest(req []byte) (resp []string) {
	resp = strings.Split(string(req), ",")
	for i := range resp {
		if resp[i] == "" {
			resp = append(resp[:i], resp[i+1:]...)
		}
	}
	return
}
