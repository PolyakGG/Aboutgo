//В аргументах Println передавались строки. Строка представляет собой
//последовательность байтов, которые обычно представляют символы
//текста. Строки можно определять прямо в программе в виде строковых
//литералов: компилятор Go интерпретирует текст, заключенный в двойные
//кавычки, как строку. "Hello, Go!"

// Некоторые управляющие символы, которые неудобно вводить с клавиатуры
//(символы новой строки, табуляции и т. д.), внутри строк могут представляться в виде служебных последовательностей: символа «обратный слеш»,
//за которым следует другой символ (или символы)

// - \n Символ новой строки
// - \t Символ табуляции
// - \" Двойная кавычка
// - \\ Обратный слеш

// Руны
// Если строки обычно используются для представления последовательностей символов, то руны
//в языке Go представляют отдельные символы.
//Строковые литералы заключаются в двойные кавычки ("), а рунные литералы записываются в одиночных кавычках (').
// В программах Go могут использоваться практически любые символы
//любых мировых языков, потому что в Go для хранения рун используется стандарт Юникод. Руны хранятся в виде числовых кодов,
//а не в виде символов; если передать руну функции fmt.Println, то выведется числовой код, а не исходный символ.
// 'A' - [65] - Выводит код символа в Юникоде.
// 'B' - [66] - Выводит код символа в Юникоде.
// 'Ж' - [1174] - Выводит код символа в Юникоде.

// Почему этот код не компилируется?
//
//
//s := "hello"
//s[0] = 'H'
//fmt.Println(s)

// Строки Go являются неизменяемыми и ведут себя как байтовые срезы только для чтения (с несколькими дополнительными свойствами).
//
// Чтобы обновить данные, используйте взамен срез рун.
//package main
//
//import "fmt"
//
//func main() {
//	buf := []rune("hello")
//	buf[0] = 'H'
//	s := string(buf)
//	fmt.Println(s) // "Hello"
//}

//Если строка содержит только символы ASCII, вы также можете использовать байтовый срез, поскольку каждый ASCII символ занимает только 1 байт.

//Строки являются неизменяемыми ( immutable ) объектами в Go.
//Это означает, что после создания строки ее содержимое не может быть изменено.
//Однако, можно создать новую строку, объединив несколько существующих строк.

// При объявлении переменной без указания значения она будет инициализирована с нулевым значением.
//Нулевым значением типа string являются пустые кавычки (""):
// var blank string
//
//fmt.Println("peace be upon you\nupon you be peace")
//fmt.Println(`strings can span multiple lines with the \n escape sequence`)
//peace be upon you
//upon you be peace
//strings can span multiple lines with the \n escape sequence

package main

import "fmt"

func main() {
	question := "¿Cómo estás?"

	for i, c := range question {
		fmt.Printf("%v %c \n", i, c)
	}
}
