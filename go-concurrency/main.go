package main

import (
	"fmt"
	"time"
)

func greet(phase string, a chan bool) {
	time.Sleep(time.Second * 2)
	fmt.Println("hello world!", phase)
	a <- true
}

func greetSlow(phase string, a chan bool) {
	time.Sleep(time.Second * 3)
	fmt.Println("Hello world, phase::: ", phase)
	a <- true
	close(a)
}

//  - main() được chạy trong main goroutine.
//   - Ba goroutine được tạo ra chạy độc lập, không “nằm bên trong” main goroutine theo kiểu cha–con.
//   - Khi main() kết thúc, process/program kết thúc.
//   - Vì process đã kết thúc nên mọi goroutine còn lại cũng bị dừng theo. Không phải main goroutine chủ động “giết” chúng.

func main() {
	// done := make(chan bool)
	// go greet("Hello nice to meet you", done)
	// go greetSlow("waiting...", done)
	// go greet("waiting for loop...", done)
	// b := <-done // done đứng chờ cho đến khi done = true, nhận tín hiệu, main bị block ở đây
	// <-done
	// <-done
	// fmt.Println(b) // log ra tín hiệu

	// cách khai báo trên phải <-done 3 lần và tính theo số lần lâu nhất mới kết thúc, khá giống async của nodejs
	// vậy để viết gọn hơn thì nên dùng make array channels

	// 	dones := make([]chan bool, 3)

	// 	dones[0] = make(chan bool)
	// 	go greet("Hello nice to meet you", dones[0])

	// 	dones[1] = make(chan bool)
	// 	go greetSlow("waiting...", dones[1])

	// 	dones[2] = make(chan bool)
	// 	go greet("waiting for loop...", dones[2])

	// 	for _, done := range dones {
	// 		<-done
	// 	}

	//cách 3.
	done := make(chan bool)
	go greet("Hello nice to meet you", done)
	go greetSlow("waiting...", done)
	go greet("waiting for loop...", done)
	
	for range done {
	}

}
