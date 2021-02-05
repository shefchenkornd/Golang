package main

import (
	"flag"
	"fmt"
)

var booleanFlag = flag.Bool("verbose", false, "help")
var msg = flag.String("msg", "default-msg", "some string")

func init()  {
	flag.Parse()
}

// go build main.go
// launch:
// ./main --msg="df" -verbose=true
// можно запустить флаг как с одним дефисом, так и с двумя!
func main() {
	fmt.Println(msg, *msg)
	fmt.Println(booleanFlag, *booleanFlag)
}
