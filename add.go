package main
import(
	"fmt"
)
func main(){
fmt.Println("hello")
fmt.Println(add(4,8))
}
func add(x int, y int) int{
	z := 0
	z = x + y
	return z
}