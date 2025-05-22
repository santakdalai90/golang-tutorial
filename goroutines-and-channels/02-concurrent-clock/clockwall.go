package main

import (
    "io"
    "log"
    "net"
    "os"
    "strings"
    "fmt"
    "strconv"
    "bufio"
    "sync"
)

type ClockConfig struct {
    Name string
    Server string
    Port int
}

func showClock(clock ClockConfig, wg *sync.WaitGroup) {
    defer wg.Done()
    log.Println("Connecting to clock server", clock.Name, clock.Server, clock.Port)
    conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", clock.Server, clock.Port))
    if err != nil {
        log.Println("error connecting to clock server",  clock.Name, err)
        return 
    }
    defer conn.Close()

    reader := bufio.NewReader(conn)
    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            if err == io.EOF {
                log.Println("connection closed by server", clock.Name)
                return
            } else {
                log.Println("error reading from clock server", clock.Name, err)
                return
            }
        }

        fmt.Printf("%-*s%s\n", 12, clock.Name+"--->", strings.TrimSpace(line))
    }
}

func main() {
    clocks := make([]ClockConfig, 0)

    for i:=1; i< len(os.Args); i++ {
        parts := strings.Split(os.Args[i], "=")
        name := parts[0]
        server := strings.Split(parts[1], ":")[0]
        port, err := strconv.Atoi(strings.Split(parts[1], ":")[1])
        if err !=nil {
            log.Println("error parsing port", name, err)
            continue
        }
        clocks = append(clocks, ClockConfig{name, server, port})
    }

    var wg sync.WaitGroup
    for _, clock := range clocks {
        wg.Add(1)
        go showClock(clock, &wg)
    }

    wg.Wait()
}