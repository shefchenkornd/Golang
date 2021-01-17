package main

func foo(n int) (r int) {
	defer func() { // этот defer запуститься
		r += n
		recover()
	}()

	var foo func()
	defer foo()
	foo = func() {  // а этот defer нет!
		r += 2
	}

	return n + 1
}

func main() {
	println(foo(3))
}
