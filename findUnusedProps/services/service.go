package service

import (
    "bufio"
    "log"
    "os"
    "strings"
    "fmt"
    "regexp"
    "io/ioutil"
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

func GetMsgProps(path string) []string{
    content, err := ioutil.ReadFile(path)
    if err != nil {
        fmt.Print(err)
    }

    strContent := string(content)

    messages, err := regexp.Compile(`#{.+}`)
    if err != nil {
        fmt.Print(err)
    }
    msgProps := messages.FindAllString(strContent, -1)

    msgPropsString := strings.Join(msgProps,"")
    filteredMessages, err := regexp.Compile(`#{\|?(\w+(\.\w+)*)`)
    if err != nil {
        fmt.Print(err)
    }
    msgPropsFiltered := filteredMessages.FindAllString(msgPropsString, -1)

    msgPropsStringFinal := strings.Join(msgPropsFiltered,"")
    filteredMessagesFinal, err := regexp.Compile(`\w+(\.\w+)*`)
    if err != nil {
        fmt.Print(err)
    }
    msgPropsFilteredFinal := filteredMessagesFinal.FindAllString(msgPropsStringFinal, -1)

    return msgPropsFilteredFinal
}

func GetUnusedKeys(msgPropsHtml []string, keys []string)(ukeys []string){
    ukeys = keys
    for i := 0; i < len(msgPropsHtml); i++{
        for j := 0; j < len(keys); j++{
            if msgPropsHtml[i] == keys[j] {
                ukeys = append(ukeys[:j], ukeys[j+1:]...)
            }
        }
    }
    return ukeys
}

func WriteSliceToFile(ukeys []string){
    fileHandle, _ := os.Create("output.txt")
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

    for i := 0; i<len(ukeys);i++ {
        fmt.Fprintln(writer, ukeys[i])  
    }	
    
    writer.Flush()
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