package main

import (
	"bufio"
	"fmt"
	"net"
)

func main()  {
	conn, _ := net.Dial("tcp", "golang.org:443")

	fmt.Println(conn, 1)

	//fmt.Fprintf(conn, "GET / HTTP/2\r\n\r\n")

	status, _ := bufio.NewReader(conn).ReadString('\n')

	fmt.Println(status)
}
