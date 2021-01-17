У программ есть несколько стандартных файловых дескрипторов

* STDIN ( os.Stdin )
* STDOUT ( os.Stdout )
* STDERR ( os.Stderr )

Сообщения для пользователя нужно выводить в os.Stdout , например через fmt.Printf .
Ошибки нужно выводить в os.Stderr , например через log.Printf .
Если ошибка фатальная, нужно прервать выполнение программы и вернуть ненулевой код выхода, например c помощью os.Exit(1) или log.Fatalf("message") или panic



### Публичные и приватные идентификаторы

Публичные идентификаторы - те, которые видны за пределами вашего пакета. Публичные идентификаторы начинаются с заглавной буквы.
Приватные идентификаторы - начинаются со строчной буквой и видны только в внутри пакета.

Структуры могут содержать как приватные, так и публичные поля
например:

```
type User struct {
    Name string         // будет видно в json.Marshal
    password string     // НЕ будет видно
}
```

### Алиасы тип данных в Go
`uintptr алиас либо на uint32 или uint64`
`byte = uint8`

`rune = uint32`

### Указатели
* `*int`
* `*string`
*  ...

у каждого типа есть свой указатель

### Литералы числовых типов
<pre>
* 42              // десятичная система
* 0755            // восьмеричная система (начинается с нуля)
* 0xDeadBeaf      // шестнадцатеричная система, hex (начинается с 0x...)

* 3.14            // с плавающей точкой
* .288
* 2.e+10

* 1+1i            // комплексные
</pre>


### Операции над числами

Все так же стандартно

```
+    sum                    integers, floats, complex values, strings
-    difference             integers, floats, complex values
*    product                integers, floats, complex values
/    quotient               integers, floats, complex values
%    remainder              integers

&    bitwise AND            integers
|    bitwise OR             integers
^    bitwise XOR            integers
&^   bit clear (AND NOT)    integers

<<   left shift             integer << unsigned integer
>>   right shift            integer >> unsigned integer
```

### Сроки в Go - это НЕИЗМЕНЯЕМАЯ последовательность БАЙТОВ (byte = uint8)

если нужно включить в строку кавычки или переносы строки - используем обратные кавычки
```
w := `hello
"cruel"
'world'`
```

### Что можно сделать со строками?

1. создавать <br />
`s := "hello \n world \u9333"`

2. получить доступ к байту(!) в строке <br />
`var c byte = s[0]`

3. получить подстроку (в байтах!) <br />
 `var s2 string = s[3:5]`
 
4. склеивать (это очень дорогая операция, старайтесь так не делать) <br />
`s2 := s + "again"`

5. узнавать длину в байтах  <br />
`length := len(s)`

---

> **латинские одна буква = 1 байт <br />**
> **кириллические одна буква = 2 байта**

```
latin := "hello"      // здесь будет 5 байт
cyrillic := "Привет!" // а здесь будет 13 байт, потому что на каждый кириллич.букву по 2 байта

объясню:
6 кирил.букв (Привет) * 2 байта = 12 байт
один восклиц.знак * 1 байта = 1 байт
итого: 13 байт
```

### Unicode in Go

Unicode - это расширенный алфавит, 
а UTF-8 - это способ его (Unicode) представлять

<pre>
- `Z` = `5A`
- `Я` = `D0` `AF`
- `♬` = `E2` `99` `AC`
</pre>

Количество символов в строке != длинна строки

s[i] - это i-ый байт, НЕ СИМВОЛ

> Символы Unicode в Go представлены с помощью типа `rune = int32`

Литералы рун выглядят так:
```
var r rune = 'Я'
var r rune = '\n'
var r rune = '\xff'     // последовательность байт
var r rune = '\u12e4'   // unicode code-point
```

Руны это целые числа поэтому их можно складывать:
```
s := "hello" + string('0' + 3) // "hello 3"
```

### Работа с UTF-8 in Go

используем пакет `unicode` и `unicode/utf8`

```
// получить первую руну из строки и ее размер в байтах
// размер в байтах - на сколько байт нужно сдвинуться чтобы получить след.символ в строке
utf8.DecodeRuneInString(s string) (r rune, size int)

// получить длину строки в рунах (руна - это один символ и не важно какой, кириллич, латинс.буквы)
// "Привет" - вернет длину строки в рунах равное шести 
RuneCountInString(s string) (n int)

// проверить валидность строки
ValidString(s string) bool
```
### Преобразовать в слайс

преобразоввать строку в слайс байтов или рун и дальше работать со слайсом
```
s := "привет"
slb := []byte(s)
slr := []rune(s)
fmt.Printf("%v \n", sl)
```
### Итерация по строке
по байтам
```
for i:=0; i < len(s); i++ {
    b := s[i]
    // i строго последовательно
    // b имеет тип byte, uint8
}
```

по рунам
```
for i, r := range s {
    // i может перепрыгивать значения 1,2,4,6,9...
    // r - имеет тип rune, int32
    fmt.Printf("i - %v, rune - %v \n", i, r)
}
```

#### В Go есть библиотека работы со строками `strings`

<pre>
// проверка наличия подстроки
fmt.Printf("Result Contains %t \n", strings.Contains(str, subStr))

// строка начинается с ?
fmt.Printf("Result hasPrefix %t \n", strings.HasPrefix(str, prefix))

// строка заканчивается на ?
fmt.Printf("Result hasSuffix %t \n", strings.HasSuffix(str, suffix))

// склейка строк
fmt.Printf("Result Join %s \n", strings.Join([]string{"я","ты","он","она","вместе дружная семья"}, ","))

// разбиение по разделителю
fmt.Printf("Result split %s \n", strings.Split(str, ""))
</pre>


## эффективная склейка строк
<pre>
var b strings.Builder
for i := 10; i >= 1; i-- {
    b.WriteString("Душа-\n")
    b.WriteRune('-')
}
fmt.Printf("%s", b.String())
</pre>

### Константы - неизменяемые значения, доступные только во время компиляции.
<pre>
const PI = 3 // принимает подходящий тип
const pi = 3.14 // строгий тип

const (
    TheA = "A"
    TheB = "B"
)

// const iota = 0 // Untyped int.
const (
    X = iota // 0
    Y		 // 1
    Z		 // 2 и тд.
)
</pre>

### Просто запомните это
на стеке дешево, на хипе дорого! 