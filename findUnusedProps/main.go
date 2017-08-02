package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

func main() {
    path := "./inputFiles/inputLines.properties"
    checkFile(path)
}

func checkFile(path string){
    file, err := os.Open(path)
    var keys []string
    if err != nil {
        log.Fatal(err)
    }
    
    defer file.Close()
    
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())

        if (line == "" || line[0] == '#') {
            continue;
        }

        key := strings.Split(line, "=")
        fmt.Println("key: %v", key)

        keys = append(keys, string(key))
    }

    //fmt.Printf("Keys: %#v", keys)


    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}