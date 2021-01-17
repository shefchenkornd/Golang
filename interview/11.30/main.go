package main

type S struct {
}

func (s S) F() {
}

type IF interface {
	F()
}

func InitPointer() *S {
	var s *S
	return s
}
func InitEfaceType() interface{} {
	var s S
	return s
}

func InitEfacePointer() interface{} {
	var s *S
	return s
}

func main() {
	println(InitPointer() == nil)      // true
	println(InitEfaceType() == nil)    // false
	println(InitEfacePointer() == nil) // false
}
