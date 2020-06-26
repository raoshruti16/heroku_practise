package main
import(
	"io"
	"os"
	"net/http"
	"strconv"
)
func main(){
	port := os.Getenv("PORT")
	http.HandleFunc("/",hi)
	http.ListenAndServe(":"+port, nil)
}
func add(x int , y int)int{
	z := 0
	z = x + y
	return z
}
func hi(w http.ResponseWriter, r *http.Request){
	s1 := add(4,8)
	s2 := strconv.Itoa(s1)
	io.WriteString(w, s2)
}
