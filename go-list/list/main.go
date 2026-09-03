package lists

import "fmt"

func main() {
	
	// unpacking list in golang = destructuring list in javascript
	list123 := []int{1,2,3}
	list456 := []int{4,5,6}
	merge := append(list123, list456...)
	fmt.Println(merge)
	
}
