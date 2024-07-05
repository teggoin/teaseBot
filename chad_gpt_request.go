package main

type ChadGptRequest struct {
	Message string      `json:"message"`
	ApiKey  interface{} `json:"api_key"`
}
