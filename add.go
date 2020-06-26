package main
import(
	"io"
	"fmt"
	"os"
	"net/http"
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
	io.WriteString(fmt.Sprint("%d\n",z)
	io.WriteString(w, "Hello World!")
}
