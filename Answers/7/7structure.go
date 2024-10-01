// В сегментах и картах хранятся значения ОДНОГО типа

// Ключевое слово «struct».

// struct {
// field1 string
// field2 int
//}

// Структура представляет собой значение, которое строится
//из других значений разных типов. Если в сегменте могут
//храниться только строковые значения, а в карте — только
//целочисленные значения, вы можете создать структуру для
//хранения строковых значений, значений int, float64, bool
//и т. д. — и все это в одной удобной группе.

//
//Тип структуры объявляется ключевым словом struct, за которым следуют фигурные скобки.
//В фигурных скобках определяются одно или несколько полей: значений, группируемых в структуре.
//Определение каждого поля размещается в отдельной строке и состоит из имени поля, за ним следует тип
//значения, которое будет храниться в этом поле.

//var myStruct struct {
//	number float64
//	word string
//	toggle bool
//}
//fmt.Printf("%#v \n", myStruct) - struct { number float64; word string; toggle bool }{number:0, word:"", toggle:false}

// Если вызвать функцию Printf с глаголом %#v, она выведет значение из myStruct в виде литерала структуры.
//Литералы структур будут рассмотрены позднее в этой главе, а пока мы видим, что полю number
//структуры присвоено значение 0, полю word — пустая строка, а полю toggle — значение false. Каждое
//поле заполняется нулевым значением своего типа.

// До настоящего момента мы использовали оператор «точка»
//для обозначения функций, «принадлежащих» другому пакету,
//или методов, «принадлежащих» некоторому значению:
// fmt.Println("hi") ; var myTime time.Time
//myTime.Year()

// Аналогичным образом оператор «точка» может использоваться для обозначения полей, «принадлежащих» структуре.
//Этот синтаксис подходит как для присваивания значений, так и для их чтения.

// myStruct.number = 3.14 - [myStruct]- Значение-Структура [.number] - Имя поля
// fmt.Println(myStruct.number) - [myStruct]- Значение-Структура [.number] - Имя поля

// Теперь мы можем воспользоваться оператором «точка» для
//присваивания значений всем полям myStruct и их последующего вывода:

//var myStruct struct {
//number float64
//word string
//toggle bool
//}
//myStruct.number = 3.14 - Присваивание значений полям структуры.
//myStruct.word = "pie" - Присваивание значений полям структуры.
//myStruct.toggle = true - Присваивание значений полям структуры.
//fmt.Println(myStruct.number) - [3.14] Чтение значений из полей структуры.
//fmt.Println(myStruct.word) - [pie] Чтение значений из полей структуры.
//fmt.Println(myStruct.toggle) - [true] Чтение значений из полей структуры.

//package main
//
//import "fmt"
//
//func main() {
//	var subscriber struct {
//		name   string
//		rate   float64
//		active bool
//	}
//	subscriber.name = "Sergey Polyakov"
//	subscriber.rate = 4.99
//	subscriber.active = true
//	fmt.Println("Name:", subscriber.name)
//	fmt.Println("Montly Rate:", subscriber.rate)
//	fmt.Println("Active? -", subscriber.active)
//}

//Структуры вроде бы удобны, но объявлять
//структурные переменные довольно утомительно. Приходится повторять все определение
//типа структуры для каждой новой переменной!

// Чтобы написать определение типа, используйте ключевое слово type, за которым следует имя
//нового определяемого типа и имя базового типа, на основе которого он должен создаваться.
//Если в качестве базового используется тип структуры, укажите ключевое слово struct
//со списком определений полей, заключенным в фигурные скобки — по аналогии с тем, как это делалось при объявлении переменных для структур.

//package main
//
//func main() {
//	type myType struct { - [type] - Ключевое слово type, [myType] - Имя определяемого типа., [struct] -  Базовый.
//	// поля}
//}

//package main
//
//import "fmt"
//
//type part struct {
//	description string
//	count       int
//}
//type car struct {
//	name     string
//	topSpeed float64
//}
//
//func main() {
//	var porsche car
//	porsche.name = "Porsche 911 R"
//	porsche.topSpeed = 323 -
//	fmt.Println("Name:", porsche.name) - - [Name: Porsche 911 R]
//	fmt.Println("Top speed:", porsche.topSpeed) - [Top speed: 323]
//
//	var bolts part
//	bolts.description = "Hex Bolt"
//	bolts.count = 24
//	fmt.Println("Description:", bolts.description) - [Description: Hex Bolt]
//	fmt.Println("Count:", bolts.count) - [Count: 24]
//}
// После того как переменные будут объявлены, мы можем
//присваивать значения полям их структур и читать их, как это
//делалось в предыдущих программах.

//теперь мы можем определить тип subscriber на уровне пакета.
//Тип структуры записывается только один раз, как базовый тип для определяемого типа.
//Когда дойдет до объявления переменных, тип структуры не нужно будет записывать снова, достаточно указать subscriber в качестве типа.
//Повторять все определение структуры уже не нужно!

// package main
//
//	type subscriber struct {
//		name   string
//		rate   float64
//		active bool
//	}
//
//	func applyDiscount(s *subscriber) {
//		s.rate = 4.99
//	}
//
//	func main() {
//		var s subscriber
//		applyDiscount(&s)
//
// }

package main

import "fmt"

type part struct {
	description string
	count       int
}

func doublePack(p *part) {
	p.count *= 2
}
func main() {
	var fuses part
	fuses.description = "Fuses"
	fuses.count = 5
	doublePack(&fuses)
	fmt.Println(fuses.description)
	fmt.Println(fuses.count)
}

// Как и в случае с сегментами и картами, Go предоставляет литералы структур для создания структур
//одновременно с инициализацией их полей.

//Синтаксис имеет много общего с литералами карт. Сначала указывается тип, за ним идут фигурные скобки.
//В фигурных скобках можно задать значения полей структуры (всех или некоторых); для каждого поля
//указывается имя, двоеточие и значение. Если вы указываете несколько полей, разделите их запятыми.

// myCar := car{name: "Corvette", topSpeed: 337} - Тип структуры[car] , Поле[name:], Значение[Corvette] , Поле[topSpeed:], Значение[337]

// Возможно, вы заметили, что в этой главе нам в основном приходилось использовать
//длинные объявления для переменных структур (если только структура не возвращалась функцией). Литералы структур позволяют использовать короткие объявления
//переменных для только что созданных структур.

// Некоторые (и даже все) поля не нужно указывать в фигурных скобках. Пропущенные
//поля инициализируются нулевыми значениями для своих типов.

// subscriber := magazine.Subscriber{Rate: 4.99}
//fmt.Println("Name:", subscriber.Name) - [Name:]
//fmt.Println("Rate:", subscriber.Rate) - [Rate: 4.99]
//fmt.Println("Active:", subscriber.Active) - [Active: false]
