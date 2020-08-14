package main

import (
    "fmt"
    "sync"
)

func main() {
    var inp int
    fmt.Scan(&inp)

    var wg sync.WaitGroup
    wg.Add(inp)

    go func(){
        
    }
}
