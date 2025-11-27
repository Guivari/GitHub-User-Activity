package githubapi

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)

func MakeRequest(user string) *http.Request {
	request := fmt.Sprintf("https://api.github.com/users/%s/events",user)
	req,err := http.NewRequest("GET",request,nil)
	if (err != nil) {
		panic(err)
	}
	req.Header.Set("User-Agent", "go-client")
	req.Header.Set("Accept", "application/json")
	return req
}

func SendRequest(req *http.Request) *http.Response {
	client := &http.Client{}
	resp, err := client.Do(req)
	if (err != nil) {
		panic(err)
	}
	return resp
}

func HandleResponse(resp *http.Response) []interface{} {
	switch resp.StatusCode {
			case 200: 
				fmt.Println("200: User found")
				break
			case 404: 
				panic("404: User not found")
	}
	data,err := io.ReadAll(resp.Body)
	if (err != nil) {
		panic(err)
	}
	var jsonData interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		panic(err)
	}
	return jsonData.([]interface{})
}