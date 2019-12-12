package _01_hello_golang

import (
	"testing"
)

func TestName(t *testing.T) {
	name := getName()
	if name != "World!" {
		t.Error("Respone from getName is unexpected value")
	}
}