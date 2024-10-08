/*
Зачем нужны дженерики в Go?

Несколько лет назад у нас в Каруне возникла необходимость перевести несколько сервисов с PHP на GO.
Как сейчас помню, в первой же программе потребовалось проверить, существует ли строка в некотором слайсе строк.
Беглый гуглёж показал, что встроенной функции, аналогичной пхпшной in_array() в языке нет. Поэтому пришлось написать свою, что-то типа такого:

func stringExistsInSlice(val string, values []string) bool {
    for _, v := range values {
        if val == v {
            return true
        }
    }

    return false
}
*/

package Generic

/*
Но проблема в том, что когда надо поискать int в слайсе интов, функция получается абсолютно такой же, отличие только в сигнатуре.

func existsInSlice(val int, values []int) bool {
    for _, v := range values {
        if val == v {
            return true
        }
    }

    return false
}

*/

/* Написать универсальную функцию под все типы — задача не очень простая.
Можно использовать reflect и interface{}, как в примере на stackoverflow, но это, понятное дело, выглядит не очень и подвержено ошибкам, не проверяемым в момент компиляции.
Или же можно использовать кодогенерацию, что тоже в общем-то так себе, так как это лишний шаг при билде.

Забегая вперёд, в go 1.18 это будет решаться так:

func existsInSlice[T comparable](val T, values []T) bool {
    for _, v := range values {
        if val == v {
            return true
        }
    }

    return false
}
*/

/*
Вот простейший пример. Функция, построчно печатающая элементы слайса любого типа.

func PrintSlice[T any](s []T) {
    for _, v := range s {
        fmt.Println(v)
    }
}
т.е. после имени функции в квадратных скобках описан некий идентификатор T (который дальше будет использован там, где вы бы раньше использовали обычный тип)
и констрейнт (ограничение) для этого типа (констрейнт any означает, что в качестве T можно передать любой тип).
Таких идентификаторов может быть несколько (через запятую); констрейнт для них указывать обязательно, это будет подробнее описано ниже.

А вот так происходит вызов такой функции, уже с конкретным типом string:
greetings := []string{"Hello", "world"};
PrintSlice[string](greetings)

Т.е. синтаксически по сути мы передаём в функцию тип как обычный аргумент (параметр).
Просто такие "аргументы" передаются и описываются в сигнатуре в отдельных квадратных скобках вместо круглых.
Поэтому функциональность так и называется: type parameters.

В некоторых случаях явным образом передавать тип не надо, компилятор его сможет вывести сам по переданным аргументам:

greetings := []string{"Hello", "world"};
PrintSlice(greetings)
*/

/*
Констрейнты (ограничения типов)
У каждого параметра-типа обязательно указывается ограничение типа.
func [T MyConstraint] (...
, где MyConstraint — это интерфейс, который описывает, каким может быть тип.
Этот интерфейс может быть обычным go-интерфейсом, описывающим требуемые методы.

type MyConstraint interface {
    String() string
}

А может быть интерфейсом, перечисляющим полный список типов, для которых он может быть использован.


type MyConstraint interface {
 int | int8 | int16 | int32 | int64
}

*/
