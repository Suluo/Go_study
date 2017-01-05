package main
import(
    "net/http"
    "os"
    "strings"
    "log"
)
type TraceHandler struct{
     h http.Handler
}
func (r TraceHandler) ServeHTTP(w http.ResponseWriter,req *http.Request){
    println("get", req.URL.Path, "from", req.RemoteAddr)
    r.h.ServeHTTP(w, req)
}

func main(){
    port := "8008"
    if len(os.Args)>1{port = strings.Join(os.Args[1:2],"")}
    h := http.FileServer(http.Dir("."))
    http.Handle("/", TraceHandler{h})
    println("Listening on port", port, "...")
    log.Fatal("ListenAndServe: ", http.ListenAndServe(":"+port, nil))
}
