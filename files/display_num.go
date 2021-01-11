package main

import "fmt"

func main() {
    for i := 1; i < 25; i++ {
        fmt.Println(i)
    }

    i := 25
    for i <= 50 {
        fmt.Println(i)
        i++
    }
}
