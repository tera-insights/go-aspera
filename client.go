package main

import (
	"context"
	"net/http"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Do(ctx context.Context, req *http.Request) {
	// do something
}

func (c *Client) NewRequest(method, relativeURL string, body interface{}) (*http.Request, error) {
	return http.NewRequest(method, relativeURL, nil)
}
