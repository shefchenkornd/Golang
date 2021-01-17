package main

import (
	"fmt"
	"strings"
)

func main() {
	// тк. строки read-only, каждая склейка через '+' или '+=' приводит к выделению памяти.
	// чтобы оптимизировать число аллокаций используйте 'strings.Builder'

	//	эффективная склейка строк
	var b strings.Builder
	for i := 10; i >= 1; i-- {
		b.WriteString("Душа-\n")
		b.WriteRune('-')
	}
	fmt.Printf("%s", b.String())
}
