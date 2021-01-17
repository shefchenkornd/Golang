
package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	// регистрируем обработчики URL-адресов
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/goodbye/", goodbye)
	http.HandleFunc("/", homePage)

	// Запускаем веб-сервер, подключенного к порту 8080
	http.ListenAndServe("localhost:8080", nil)
}

func hello(res http.ResponseWriter, req *http.Request)  {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Incognito"
	}

	fmt.Fprint(res, "Hello, my name is ", name)
}

func goodbye(res http.ResponseWriter, req *http.Request)  {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	name := parts[2]
	if name == "" {
		name = "Incognito"
	}
	fmt.Fprint(res, "Goodbye ", name)
}


func homePage(res http.ResponseWriter, req *http.Request)  {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "The homepage!")
}