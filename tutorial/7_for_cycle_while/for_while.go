package main

//the for loop in go can be used as while
import "fmt"

func main() {
    sum:=1
    //while sum is less than 1k 
    for sum<1000 {
        sum+=sum
    }
    fmt.Println(sum)
}
