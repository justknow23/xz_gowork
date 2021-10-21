package zip

import "strings"

type GreetingRequest struct {
	Name string
}
type GreetingResponse struct {
	Status int
	Data   string
}

func Greeting(req *GreetingRequest) *GreetingResponse {
	response := &GreetingResponse{}
	dataWriter := &strings.Builder{}
	dataWriter.WriteString("hello,")
	if len(req.Name) == 0 {
		req.Name = "word"
	}
	dataWriter.WriteString(req.Name)
	response.Data = dataWriter.String()
	response.Status = 200
	return response
}
