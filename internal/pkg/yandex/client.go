package yandex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	url    string
	apiKey string
}

func NewClient(url string, apiKey string) *Client {
	return &Client{url: url, apiKey: apiKey}
}

func (c Client) Do(data Translate) (*TranslationsResult, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	bodyReader := bytes.NewReader(b)
	req, err := http.NewRequest(http.MethodPost, c.url, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "Application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Api-Key %s", c.apiKey))

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("yandex client error: status is %q", res.Status)
	}

	resBody, err := ioutil.ReadAll(res.Body)

	var result *TranslationsResult
	err = json.Unmarshal(resBody, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
