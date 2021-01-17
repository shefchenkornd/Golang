package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

// go tool pprof -svg -aloc_objects 0. pprof.exe mem_out.txt > mem_ao.svg
// go tool pproff -svg 0.pprof.exe cpu_out.txt > cpu.svg


type Post struct {
	ID       int
	Text     string
	Author   string
	Comments int
	Time     time.Time
}

func NewPost(ID int) *Post {
	return &Post{ID: ID}
}
// handle Post ResponseWriter
func handle(w http.ResponseWriter, req *http.Request) {
	s := ""
	for i := 0; i < 1000; i++ {
		p := &Post{ID: i, Text: "new post"}
		s += fmt.Sprintf("%#v", p)
	}

	w.Write([]byte(s))
}

type message struct {
	Field1 int 	  `json:"field_1"`
	Field2 string `json:"field_2"`
}

func main() {
	msg := message{
		Field1: 1,
		Field2: "2",
	}




	http.HandleFunc("/", handle)

	fmt.Println("starting server at :8080")

	http.ListenAndServe("localhost:8080", nil)
}

