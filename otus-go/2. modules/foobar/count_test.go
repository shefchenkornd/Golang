package foobar

import "testing"

func TestCount(t *testing.T) {
	s := "qwerasdfe"
	e := 2
	if c := Count(s, 'e'); c != e {
		t.Fatalf("bad count for %s: got %d expected %d", s, c, e)
	}
}

// Важно:
//* `package`  тот же что и в основном коде
//* Все тесты - экспортируемые функции, начинающиеся с букв `Test`
//* Тесты принимают `t` - API для установки результата тестирования