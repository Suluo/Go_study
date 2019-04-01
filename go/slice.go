package main
import "fmt"
func main(){
    var a = make([]int, 3, 3)
    for i:=0;i<3;i++ {
        a[i] = i
    }
    fmt.Println(a,&a)
    s := fmt.Sprintf("%p", &a)
    fmt.Println(s)
    a = append(a, 5)
    fmt.Println(a,&a)
    s = fmt.Sprintf("%p", &a)
    fmt.Println(s)
}
