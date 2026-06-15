package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

func main() { // 1 goroutine
	// howManyGoRoutine() // -> 2 goroutines.  Có thể chỉ cần 1 OS thread.
	// moreGoRoutines()   // -> 11 goroutines. Có thể chỉ cần 1 OS thread.
	// checkGomaxProcs()
	// cpuBound()
	f, _ := os.Create("trace.out")
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()

	fmt.Printf("GOMAXPROCS hiện tại: %d\n", runtime.GOMAXPROCS(0))
	testWork5m()
}

func worker() {
	fmt.Println("Worker is working...")
}

func howManyGoRoutine() { // 1 goroutine
	// go-routine
	// example 1: define how many go-routines
	go worker()
	time.Sleep(time.Second) // là thời gian để go-routine hoàn thành công việc trước khi main function kết thúc
}

// ex2: define number of go-routines and processer running
func moreGoRoutines() {
	for i := 0; i < 10; i++ {
		go func(id int) {
			fmt.Println(id)
		}(i)
	}
	time.Sleep(time.Second)
	//Worker is working... G >>> P >>> M
	// 9
	// 6
	// 7
	// 8
	// 3
	// 2
	// 4
	// 0
	// 1
	// 5
}

// GOMAXPROCS là số luồng tối đa được phép sử dụng để chạy các goroutine cùng một lúc
func checkGomaxProcs() {
	for i := 0; i < 100; i++ {
		go func(id int) {
			for {
			}
		}(i)
	}
	// Đặt GOMAXPROCS thành 4
	runtime.GOMAXPROCS(2)

	// In ra số lượng đang được áp dụng
	fmt.Println("GOMAXPROCS hiện tại:", runtime.GOMAXPROCS(0))
}

func work(_ int) {
	for {
	}
}
func cpuBound() {
	runtime.GOMAXPROCS(4)
	for i := 0; i < 1000; i++ {
		go work(i)
	}
	// 1000 G
	// 4 P
	// ≈ 4 M
}

func work5m(_ int, wg *sync.WaitGroup) {
	defer wg.Done() // Báo cáo đã hoàn thành khi thoát hàm

	// Giả lập xử lý: Thay bằng time.Sleep (I/O Bound) hoặc tính toán (CPU Bound)
	time.Sleep(5 * time.Millisecond)
}
func testWork5m() {

	var wg sync.WaitGroup
	totalGoroutines := 100000

	startTime := time.Now()

	for i := 0; i < totalGoroutines; i++ {
		wg.Add(1)
		go work5m(i, &wg)
	}
	fmt.Printf("Số lượng Goroutines đang hoạt động: %d\n", runtime.NumGoroutine())
	wg.Wait()
	fmt.Printf("Hoàn thành! Tổng thời gian chạy: %v\n", time.Since(startTime))
}
