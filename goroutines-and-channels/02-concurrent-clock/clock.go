package main

import (
    "io"
    "log"
    "net"
    "time"
    "flag"
    "fmt"
)

func handleConn(c net.Conn, tz *string) {
    defer c.Close()

    location, err := time.LoadLocation(*tz)
    if err != nil {
        panic(err)
    }

    for {
        _, err := io.WriteString(c, time.Now().In(location).Format("15:04:05\n"))
        if err != nil {
            log.Println("Client disconnected", err)
            return
        }
        time.Sleep(1* time.Second)
    }
}

func main() {
    port := flag.Int("port", 8080, "port to listen on")
    tz := flag.String("tz", "Asia/Kolkata", "timezone to use")
    flag.Parse()
    
    listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Clock Server started on localhost:%d. Use nc localhost %d to connect", *port, *port)

    for {
        conn, err := listener.Accept()
        log.Println("Received an incoming connection")
        if err != nil {
            log.Println(err)
            continue
        }
        go handleConn(conn, tz)
    }
}