package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := []byte(`
	{
		"ok": true,
		"total": 3,
		"documents": [
			{"id":11, "title":"c++"},
			{"id":2, "title":"suddenly perl"},
			{"id":5, "title":"go"}
		]
	}
	`)

	var resp struct {
		Ok        bool `json:"ok"`
		Total     int  `json:"total"`
		Documents []struct {
			id    int    `json:"id"`
			Title string `json:"title"`
		} `json:"documents"`
	}

	err := json.Unmarshal(data, &resp) // указатель на структуру куда мы всё это размаршалим
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
	fmt.Println(resp.Documents[0].Title)
}
