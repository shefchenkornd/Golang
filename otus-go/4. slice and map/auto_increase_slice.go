package main

import "fmt"

/**
Если len < cap - увеличивается len.
Если len = cap - то увеличивется cap, выделяется новый кусок памяти, данные копируются!
 */
func main() {
	sl := []int{1}

	for i := 0; i < 100; i++ {
		fmt.Printf("len: %d \tcap: %d \tptr %0x\n",
			len(sl), cap(sl), &sl[0])

		sl = append(sl, i)
	}
}
