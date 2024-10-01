// В Go (Golang) виртуальные потоки, управляемые на уровне приложения, называются горутинами (goroutines).
// Горутины — это легковесные потоки, которые управляются рантаймом Go, а не операционной системой.
// Это позволяет создавать и управлять тысячами горутин с минимальными затратами по сравнению с традиционными потоками операционной системы.
//
// Основные особенности горутин:
// Легковесность: Горутины потребляют значительно меньше памяти и ресурсов по сравнению с потоками операционной системы.
// Они стартуют с небольшим стеком (около 2Models.queue КБ), который динамически растет и уменьшается по мере необходимости.
// Планировщик Go: Горутины управляются планировщиком Go, который распределяет их выполнение между доступными потоками операционной системы.
// Это позволяет эффективно использовать многопроцессорные системы.
// Простота использования: Создание горутины в Go очень просто и не требует сложной настройки.
// Для этого достаточно использовать ключевое слово go перед вызовом функции.
// Коммуникация через каналы: Горутины могут общаться друг с другом с помощью каналов (channels), что обеспечивает безопасную передачу данных между ними без необходимости использования мьютексов или других механизмов синхронизации.
// Пример использования горутин и каналов:
package main

import (
	"fmt"
	"time"
)

// Функция, которая будет выполняться в горутине

func worker(id int, ch chan string) {
	time.Sleep(time.Second)
	ch <- fmt.Sprintf("Worker %d done", id)
}

func main() {
	// Создаем канал для передачи строк
	ch := make(chan string)

	// Запускаем несколько горутин
	for i := 1; i <= 5; i++ {
		go worker(i, ch)
	}

	// Получаем результаты от горутин
	for i := 1; i <= 5; i++ {
		fmt.Println(<-ch)
	}
}

//Объяснение примера:
//Функция worker: Эта функция принимает идентификатор и канал. Она имитирует выполнение некоторой работы с помощью time.Sleep и затем отправляет сообщение в канал.
//Создание канала: В функции main создается канал для передачи строк.
//Запуск горутин: В цикле запускаются пять горутин, каждая из которых выполняет функцию worker.
//Получение результатов: В другом цикле основной поток получает результаты от горутин через канал и выводит их на экран.
//Преимущества горутин:
//Высокая производительность: Благодаря легковесности и эффективному управлению, горутины позволяют создавать высокопроизводительные приложения.
//Простота параллелизма: Использование горутин и каналов упрощает написание параллельных и конкурентных программ.
//Масштабируемость: Приложения на Go легко масштабируются благодаря возможности создания большого количества горутин.
//Горутины являются одной из ключевых особенностей Go, делающих этот язык привлекательным для разработки высокопроизводительных и масштабируемых приложений.
