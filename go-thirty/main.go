
//Step 1: Execute go mod init <module-name>
//Step 2: Execute go get github.com/gruntwork-io/terratest/modules/random
package main
import (
	"fmt"
	"github.com/gruntwork-io/terratest/modules/random"
)

func main() {
	fmt.Println("Hello go")
	a := random.RandomInt([]int{1, 2, 3, 4, 5})
	fmt.Printf("Random int: %d\n", a)
}
