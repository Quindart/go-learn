// Bài 1.1 (WaitGroup): Viết chương trình chạy 5 Goroutine song song,
//  mỗi Goroutine in ra số thứ tự của nó (từ 1 đến 5).
//  Dùng sync.WaitGroup để đảm bảo hàm main chờ cả 5 Goroutine chạy xong mới kết thúc.

// Bài 1.2 (Channel): Tạo một Unbuffered Channel.
//  Cho 1 Goroutine gửi chuỗi "Hello Fresher" vào channel,
//  và hàm main nhận giá trị đó rồi in ra màn hình.

// Bài 1.3 (Select): Tạo 2 channel ch1 và ch2.
// Dùng select để in ra dữ liệu từ channel nào phản hồi trước.

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 1.1 WaitGroup
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d done\n", id)
		}(i)
	}
	wg.Wait()

	// 1.2 Channel
	ch := make(chan string)
	go func() {
		ch <- "Hello Fresher"
	}()
	fmt.Println(<-ch)

	// 1.3 Select
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() { time.Sleep(100 * time.Millisecond); ch1 <- "Từ ch1" }()
	go func() { time.Sleep(50 * time.Millisecond); ch2 <- "Từ ch2" }()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
}
