package main

import (
	"bytes"
	"strings"
)

/**
интерфейсы io.Reader and io.Writer  могут быть
реализованы различными структурами в памяти
 */
func main() {
 strings.Reader{} 	// реализует io.Reader
 strings.Builder{}	// реализует io.Writer

 bytes.Reader{}		// реализует io.Reader
 bytes.Buffer{}		// реализует io.Writer, io.Reader
}
