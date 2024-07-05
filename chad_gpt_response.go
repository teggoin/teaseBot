package main

type ChadGptResponse struct {
	IsSuccess       bool   `json:"is_success"`
	Response        string `json:"response"`
	UsedWordsCount  int    `json:"used_words_count"`
	UsedTokensCount int    `json:"used_tokens_count"`
}
