package main
import "fmt"
func main(){
    result1:=0
    result2:=0
    var num int
    fmt.Scan(&num)
    for a:=num;a>0;a-- {
        result1 += a * a
    }
    result2 = num*(num+1)/2
    fmt.Println("立方体の数:", result1, "個")
    fmt.Println("立方体の面積:", result2, "cm^2")
}
