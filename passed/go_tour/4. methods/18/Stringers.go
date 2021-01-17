package main

import (
	"fmt"
	"unicode/utf8"
)

// Упражнение: Stringers
// Сделайте так, чтобы тип IPAddr реализовывал fmt.Stringer, чтобы выводить адрес как четыре значения, разделенные точкой.
// Например, IPAddr{1, 2, 3, 4} должен быть напечатан как "1.2.3.4".

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ipAddr IPAddr) String() string {
	str := ""
	for _, ip := range ipAddr {
		str += fmt.Sprintf(".%d", ip)
	}
	return trimFirstRune(str)
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
