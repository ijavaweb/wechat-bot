package model


type OpenAIResponse struct {
	Choices []Choice `json:"choices"`
}
type Choice struct {
	Message M `json:"message"`
	FinishReason string `json:"finish_reason"`
	Index int `json:"index"`
}
type M struct {
	Role string `json:"role"`
	Content string `json:"content"`
}
type OpenAIRequest struct {
	Model string `json:"model"`
	Messages []Message `json:"messages"`
}
type  Message struct {
	Role string `json:"role"`
	Content string `json:"content"`
}

type  ReqMessage struct {
	Content string `json:"content"`
}