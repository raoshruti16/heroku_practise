package main
import(
	"fmt"
	"io"
	"net/http"
)
func main(){
fmt.Println("hello")
fmt.Println(add(4,8))
port := os.Getenv("PORT")
http.HandleFunc("/",add)
http.ListenAndServe(":"+port, nil)
}
func add(x int, y int) int{
	z := 0
	z = x + y
	return z
}
