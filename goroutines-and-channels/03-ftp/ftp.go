package main

import (
    "os"
    "fmt"
    "log"
)

func listFiles() {
    files, err := os.ReadDir(".")
    if err != nil {
        log.Println(err)
        return
    }

    for _, file := range files {
        fmt.Println(file.Name())
    } 
}

func changeDirectory() {
    fmt.Println("change directory")
}

func main() {
    for {
        fmt.Print("ftp> ")
        var command string
        fmt.Scanln(&command)
        switch command {
        case "":
            continue
        case "exit":
            return
        case "ls":
            listFiles()
        case "cd":
            changeDirectory()
        default:
            fmt.Println("Unknown command")
        }
    }
}