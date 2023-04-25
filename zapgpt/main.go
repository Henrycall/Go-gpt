package main

import (
	"encoding/json"
	"fmt"
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
	}

func main(){
	fmt.Println("hello word")
}