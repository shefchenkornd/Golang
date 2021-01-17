package foobar

func Count(s string, r rune) int {
	var cnt int
	for _, l := range s {
		if l == r {
			cnt +=1
		}
	}

	return cnt
}

func main()  {
	count := Count()
}