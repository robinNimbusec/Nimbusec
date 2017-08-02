package service

import (
    "bufio"
    "log"
    "os"
    "strings"
)

func GetKeys(path string)(keys []string){ 
    file, err := os.Open(path)
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
        keys = append(keys, key[0])
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return keys
}

func GetUnusedKeys(path string, keys []string)(ukeys []string){
    file, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }
    
    defer file.Close()

    ukeys = keys
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()

        for index, key := range keys {
            if strings.Contains(line, key) {
                ukeys = append(ukeys[:index], ukeys[index+1:]...)
            }
        }
    }
    return ukeys
}

/*

// n = 10
var n []int = {1, 2, 3, 4, 5}
n[1] = 2
n[0] = 1

n[1:2] = [2, 3]
n[:4] = [1, 2, 3, 4, 5]
n[3:] = [4, 5]
n[3:]... = 4, 5

// [1, 2] 3 [4, 5]
n = append(n[:1], n[2:]...)

*/