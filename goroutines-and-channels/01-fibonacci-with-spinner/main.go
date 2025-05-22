package main

import (
    "time"
    "fmt"
)

func spinner(delay time.Duration) {
    for {
        for _, r := range `◐◓◑◒` {
            fmt.Printf("\r%c", r)
            time.Sleep(delay)
        }
    }
}

func fib(x int) int {
    if x < 2 {
        return x
    }

    return fib(x-1) + fib(x-2)
}
func main() {
    go spinner(100 * time.Millisecond)
    fmt.Printf("\rfibonacci(45) = %d\n", fib(45))
}