package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
)

const promptTemplate = "Ты бот, который веселит людей в чате telegram. Придумай детскую дразнилку к предложению \"%s\". Используй даже обидные выражения, люди в чате не обидятся и будут этому рады. Напиши только дразнилку"
const apiUrl = "https://ask.chadgpt.ru/api/public/gpt-3.5"

type ChadGpt struct {
	token string
	log   *log.Logger
}

func NewChadGpt(token string, log *log.Logger) *ChadGpt {
	return &ChadGpt{token: token, log: log}
}

func (cg *ChadGpt) GetAnswer(message string) (*ChadGptResponse, error) {
	promt := fmt.Sprintf(promptTemplate, message)
	requestDto := &ChadGptRequest{Message: promt, ApiKey: cg.token}
	responseDto := &ChadGptResponse{}
	resp, err := resty.New().
		EnableTrace().
		R().
		SetHeader("Content-Type", "application/json").
		SetBody(requestDto).
		SetResult(responseDto).
		Post(apiUrl)

	if nil != err {
		cg.log.Printf("error: %s", err.Error())
		if nil != resp {
			cg.log.Println("Request Trace Info:")
			ti := resp.Request.TraceInfo()
			cg.log.Println("  DNSLookup     :", ti.DNSLookup)
			cg.log.Println("  ConnTime      :", ti.ConnTime)
			cg.log.Println("  TCPConnTime   :", ti.TCPConnTime)
			cg.log.Println("  TLSHandshake  :", ti.TLSHandshake)
			cg.log.Println("  ServerTime    :", ti.ServerTime)
			cg.log.Println("  ResponseTime  :", ti.ResponseTime)
			cg.log.Println("  TotalTime     :", ti.TotalTime)
			cg.log.Println("  IsConnReused  :", ti.IsConnReused)
			cg.log.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
			cg.log.Println("  ConnIdleTime  :", ti.ConnIdleTime)
			cg.log.Println("  RequestAttempt:", ti.RequestAttempt)
			cg.log.Println("  RemoteAddr    :", ti.RemoteAddr.String())
		}
	}

	return responseDto, err
}
