package ddrago

import "net/http"

type Client struct {
	client http.Client
}

func New(client http.Client) *Client {
	return &Client{client:client}
}
