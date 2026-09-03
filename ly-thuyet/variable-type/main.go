package main

import "fmt"

// Go là ngôn ngữ tĩnh (statically typed),

// nghĩa là mọi biến đều phải có một kiểu dữ liệu xác định rõ ràng:
// Số nguyên (Integers): Gồm int, int8, int16, int32, int8,
// và các loại không âm uint, uint8, uint16, uint32, uint64, uintptr.

// Kiểu int mặc định thường là 32-bit hoặc 64-bit tùy theo kiến trúc máy tính.
// Số thực (Floating-point): Gồm float32 và float64 dùng để lưu số có phần thập phân.
// Số phức (Complex): Gồm complex64 và complex128.

// Boolean: Kiểu bool chỉ nhận hai giá trị là true hoặc false.

// Chuỗi (String): Kiểu string dùng để lưu trữ văn bản, được đặt trong dấu ngoặc kép "".

// Kiểu dữ liệu nâng cao/phức hợp:
//  Gồm mảng (array), lát cắt (slice), ánh xạ (map), cấu trúc (struct), và con trỏ (pointer)

func main(){

	var a int = 10
	var b int32 = 20
	//var c = a + b //invalid operation: a + b (mismatched types int and int32)compilerMismatchedTypes
	var c = a + int(b) // ép kiểu b sang int để thực hiện phép cộng
	var d float64 = 3.14
	var e = float64(a) + d // ép kiểu a sang float64 để thực hiện phép cộng

	// in ra giá trị và kiểu dữ liệu của các biến
	fmt.Println("a:", a, "type:", "int")
	fmt.Println("b:", b, "type:", "int32")
	fmt.Println("c:", c, "type:", "int")
	fmt.Println("d:", d, "type:", "float64")
	fmt.Println("e:", e, "type:", "float64")

	fmt.Println(5%3) // 2 -> chia lay dư
	fmt.Println(5.0/3.0) // 1.6666666666666667 -> chia lay thuong

	// 
	if a == 2 || a==3 {
		fmt.Println("a bằng 2 hoặc a bằng 3")
	}
	ok := true
	if ok {
		fmt.Println("ok là true")
	} else {
		fmt.Println("ok là false")
	}
	// var check bool
	// e = bool(1) // ép kiểu 1 sang bool, nhưng Go không cho phép ép kiểu từ số sang bool trực tiếp

	// strings
	s1 :="Hello"
	s2 := "world"
	s := "row1 \n row2"
	fmt.Println(s1, s2, s)

	// length of string
	fmt.Println(len(s1)) // 5

	// concat: ket noi chuoi
	s3 :=  "s3" +" " + s1 + " " + s2
	fmt.Println(s3) // s3 Hello world

	// slice chuoi
	fmt.Println(s3[0:8]) // s3 Hello
	fmt.Println(s3[3:]) // Hello world
}
