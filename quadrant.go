package main

import "fmt"

func main() {
    var x int
    var y int
    fmt.Scanf("%d", &x)
    fmt.Scanf("%d", &y)

    if x > 0 && y > 0 {
        fmt.Println(1)
    } else if x < 0 && y > 0 {
        fmt.Print(2)
    } else if x < 0 && y < 0 {
        fmt.Print(3)
    } else {
        fmt.Print(4)
    }
}