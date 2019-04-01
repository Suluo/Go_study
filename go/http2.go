package main
import(
    "fmt"
    "net/http"
    "log"
)
type Counter struct{
    n int
}
func (ctr *Counter) ServeHTTP(c http.ResponseWriter, req *http.Request){
    ctr.n++
    fmt.Fprintf(c, "counter = %d\n", ctr.n)
}

func main(){
    http.Handle("/counter", new(Counter))
    log.Fatal("ListerAndServe:", http.ListenAndServe(":12345", nil))
}
