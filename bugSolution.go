```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        ch := make(chan int)

        for i := 0; i < 5; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        fmt.Printf("Goroutine %d received: %d\n", i, <-ch)
                }(i)
        }

        wg.Add(1) // Add a goroutine to send data
        go func() {
                defer wg.Done()
                ch <- 10 // Sending data to the channel
                close(ch)
        }()

        wg.Wait()
}
```