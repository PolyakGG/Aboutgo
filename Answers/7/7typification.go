// Структурная типизация в Go
//В Go типы определяются не по названию, а по структуре (полям и методам).
//Два типа считаются совместимыми, если они имеют одинаковую структуру, независимо от их названий.
//Это позволяет использовать значения одного типа там, где ожидается другой, но совместимый тип.

// type Person struct {
//    Name string
//    Age  int
//}
//
//type Employee struct {
//    Name string
//    Age  int
//    Salary float64
//}
//
//func DescribePerson(p Person) {
//    fmt.Printf("%s is %d years old\n", p.Name, p.Age)
//}
//
//func main() {
//    p := Person{"Alice", 30}
//    e := Employee{"Bob", 35, 50000.0}
//
//    DescribePerson(p)  // Работает, так как Person и Employee имеют совместимую структуру
//    DescribePerson(e)
//}
// «Утиная» типизация в Go
//Хотя в Go нет явной поддержки «утиной» типизации, как в динамически типизированных языках, структурная типизация обеспечивает схожее поведение.
//Если тип реализует все необходимые методы интерфейса, то он считается совместимым с этим интерфейсом, независимо от названия типа.

package main

import "fmt"

type Walker interface {
	Walk()
}

type Person struct {
	Name string
}

func (p Person) Walk() {
	fmt.Printf("%s is walking\n", p.Name)
}

type Employee struct {
	Name   string
	Salary float64
}

func (e Employee) Walk() {
	fmt.Printf("%s is walking\n", e.Name)
}

func MakeWalk(w Walker) {
	w.Walk()
}

func main() {
	p := Person{"Alice"}
	e := Employee{"Bob", 50000.0}

	MakeWalk(p) // Работает, так как Person реализует интерфейс Walker
	MakeWalk(e) // Также работает, так как Employee реализует интерфейс Walker
}
