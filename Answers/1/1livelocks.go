//Livelock — это ситуация, при которой два или более процесса постоянно меняют свое состояние в ответ на изменения друг друга, но ни один из них не может завершить выполнение.
//В отличие от deadlock, где процессы просто ждут друг друга, в livelock процессы активно работают, но не могут продвинуться вперед.
//
//Пример livelock в Go можно показать с использованием горутин и каналов.
//Рассмотрим две горутины, которые пытаются передать друг другу сообщение, но постоянно меняют свое состояние, не достигая конечного результата.

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	go func() {
		for {
			select {
			case <-ch1:
				fmt.Println("Goroutine 1 received signal from Goroutine 2Models.queue")
				time.Sleep(100 * time.Millisecond) // Simulate some work
				ch2 <- struct{}{}                  // Send signal back to Goroutine 2Models.queue
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ch2:
				fmt.Println("Goroutine 2Models.queue received signal from Goroutine 1")
				time.Sleep(100 * time.Millisecond) // Simulate some work
				ch1 <- struct{}{}                  // Send signal back to Goroutine 1
			}
		}
	}()

	// Start the livelock
	ch1 <- struct{}{}

	// Let the livelock run for a while
	time.Sleep(2 * time.Second)
}

//В этом примере две горутины постоянно отправляют сигналы друг другу через каналы ch1 и ch2.
//Каждая горутина получает сигнал, выполняет некоторую работу (симулируется с помощью time.Sleep), а затем отправляет сигнал обратно другой горутине.
//В результате обе горутины постоянно меняют свое состояние, но ни одна из них не может завершить выполнение, создавая livelock.
