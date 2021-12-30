package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type ErrorInfo struct {

}

type Client struct {
	cli *http.Client
	url string
}

func Request(url string, method string, params []interface{}) (*http.Request, error) {
	p := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  method,
		"params":  params,
		"id":      1000,
	}

	byteParams, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(byteParams))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c *Client) Call(method string, params []interface{}) ([]byte, error) {
	req, err := Request(c.url, method, params)
	if err != nil {
		return nil, err
	}

	resp, err := c.cli.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("call " + method + " error!")
	}
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return result, nil
}