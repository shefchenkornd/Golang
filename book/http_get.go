package book

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("https://ya.ru/")
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	resp.Body.Close()
}