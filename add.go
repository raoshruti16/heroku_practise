package main
import(
	"io"
	"os"
	"net/http"
	"strconv"
)
func main(){
	port := os.Getenv("PORT")
	http.HandleFunc("/",add)
	http.ListenAndServe(":"+port, nil)
}
func add(w http.ResponseWriter, r *http.Request){
	x , y := 6,100
	z := 0
	z = x + y
	s2 := strconv.Itoa(z)
	io.WriteString(w, s2)
	io.WriteString(w, "Hello World!")
}
