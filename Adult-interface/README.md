# Реализовать интерфейс Adult
### Интерфейс - контракт. Для реализации интерфейса необходимо, чтобы были реализованы все его методы.

__Изначальный код:__<br>
```
package main

import (
	"fmt"
)

type Adult interface {
	IsAdult() bool
	fmt.Stringer
}

type Person struct {
	age  int
	name string
}

func adultFilter(people []Adult) []Adult {
	adults := make([]Adult, 0)
	for _, p := range people {
		if p.IsAdult() {
			adults = append(adults, p)
		}
	}
	return adults
}

func main() {
	people := []Adult{Person{15, "John"}, Person{18, "Joe"}, Person{45, "Mary"}}
	fmt.Println(adultFilter(people))
}
```
<br>

__Решение задачи:__

<br>

Была добавлена функция реализации метода `IsAdult()`


```
func (p Person) IsAdult() bool {
	if p.age >= 18 {
		return true
	} else {
		return false
	}
}
```

Была добавлена функция реализации метода `String()`


```
func (p Person) String() string {
	return fmt.Sprintf("%v years %v", p.age, p.name)
}
```

<br>

## ТЕОРИЯ

- это набор сигнатур методов
- интерфейс реализуется неявно
- интерфейс может встраивать другие интерфейсы
- имена методов интерфейса не должны повторяться
- интерфейс может быть пустым (не иметь методов), такому интерфейсу удовлетворяет любой тип

<br>

### Интерфейсы изнутри

На этапе компиляции:
- генерируются метаданные для каждого статического типа, включая его список методов
- генерируются метаданные для каждого интерфейса, включая его список методов


И при компиляции и в рантайме в зависимости от выражения:
- сравниваются methodset'ы типа и интерфейса
- создается и кэшируется `itab`

```
var s Speaker = string("test") // compile-time error
var s Speaker = io.Reader // compile time error
var h string = Human{} // compile time error

// runtime error
var s interface{};
h = s.(Human)
```