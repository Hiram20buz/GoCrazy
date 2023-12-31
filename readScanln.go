package main

import (
    "fmt"
    "log"
)

func main() {
    fmt.Println("input text:")
    var w1, w2, w3 string
    n, err := fmt.Scanln(&w1, &w2, &w3)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("number of items read: %d\n", n)
    fmt.Printf("read line: %s %s %s-\n", w1, w2, w3)
}
