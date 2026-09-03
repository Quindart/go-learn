
package main

import "fmt"

func main(){
	fmt.Println("Welcome back!")

	websites := map[string]string{
		"google": "https://www.google.com",
		"yahoo": "https://www.yahoo.com",
		"bing": "https://www.bing.com",
	}
	fmt.Println("All websites",websites)
	// mutating the map
	websites["duckduckgo"] = "https://www.duckduckgo.com"
	fmt.Println(websites)
	delete(websites, "yahoo")
	fmt.Println("Deleted websites:::",websites)
}
// khác nhau giữa map và struct là map có thể thay đổi được còn struct thì không thể thay đổi được.
// map là một tập hợp các cặp key-value, trong khi struct là một tập hợp các trường dữ liệu có kiểu dữ liệu khác nhau. 
// Map có thể được thay đổi bằng cách thêm hoặc xóa các cặp key-value, 
// trong khi struct không thể thay đổi sau khi được tạo ra.