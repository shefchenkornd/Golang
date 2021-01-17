package main

func main() {
	const (
		_ = iota
		b = iota
		kb = 1024 * iota  // здесь должно быть возведение в степень
		mb				// mb
		gb				// gb
		tb				// tb
	)
}
