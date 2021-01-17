package main

import (
	"sync"
)

type Countters struct {
	mx sync.RWMutex
	m map[string]int
}

func (c *Countters) Load(key string) (int, bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	val, ok := c.m[key]
	return val, ok
}

func (c *Countters) Store(key string, value int)  {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key] = value
}

/**
defer имеет небольшой оверхед (порядка 50-100 наносекунд), поэтому если у вас код для высоконагруженной системы
и 100 наносекунд имеют значения, то вам может быть выгодней не использовать defer

Методы Load() и Store() должны быть определены для УКАЗАТЕЛЯ на Countters,
потому что у нас должен быть указатель на мьютекс в нашей структуре.
 */
func main() {
	var wg sync.WaitGroup
	m := Countters{
		m: map[string]int{},
	}

	/**
	fatal error: concurrent map writes
	т.е. одновременная запись map!! такая ошибка вас ждет)
	ибо писать одновременно в map нельзя!!!
	*/
	for x := 0; x < 12; x++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			m.Store("hello", 1)
		}(&wg)
	}

	wg.Wait()
}
