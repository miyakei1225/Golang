package main

import "fmt"

func main() {
    for i := 1; i <= 50; i++ {
        fmt.Print(i)
        if i % 3 == 0 {
            fmt.Print("プロサー")
        }
        fmt.Print("\n")
    }
}
