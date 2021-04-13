package api

import (
	"net/http"
)

func (c *Client) PostValidationCode() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		http.Error(writer, "", 200)
	}
}

func (c *Client) PostToken() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
	}
}
