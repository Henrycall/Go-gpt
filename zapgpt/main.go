package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)


type Message struct {
	Model string `json:"role"`
	Messages []Message `json:"content"`
}

type Request struct {
	Role string `json:"model"`
	Content string `json:"messages"`
	MaxTokenes int `json:"max_tokens"`
}

type Response struct {
	ID string `json:"id"`
	Object string `json:"object"`
	Created int `json:"created"`
	Choices []Choices `json:choices`
}
type Choices struct {
	Index int `json:"index"`
	Message struct {
		Role string `json:"role"`
		Contet string `json:"content"`
	} 
}

	func genereteGPT(query string) (string,error){
		req :=Request{
			Model: "gpt-3.5-turbo",
			Messages: []Message{
				{
					Role: "user",
					Content: query,
				},
			},
			MaxTokenes: 300,
		}
		
		reqToJson, err := json.Marshal(req)
		if err != nil {
			return "", err
		}
		request,err :=  http.NewRequest("POST","https://api.openai.com/v1/chat/completions", bytes.NewBuffer(reqToJson))
		if err != nil {
			return "", err
		}
		request.Header.Set("Content-Type" , "application/json")
		request.Header.Set("Authorization" , "Bearer sk-xxx")

		response,err = http.DefaultClient.Do(request)
		if err != nil {
			return "",err
		}
		defer response.Body.Close()
		respBody,err := ioutil.ReadAll(response.Body)
		var resp Response
		err = json.Unmarshal(response.Body ,&resp )
		if err != nil {
			return "",err
		}
		return resp.Choices[0].Message.Contet , nil
	}


func main(){
	fmt.Println("hello word")
}