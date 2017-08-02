package main 

import (
    "fmt"
    "../findUnusedProps/services"
    "os"
)

func main() {
    pathToInputLines := os.Args[1]//"./inputFiles/inputLines.properties"
    pathForSearching := os.Args[2]//"./inputFiles/research.html"

    //fmt.Prinln()
    keys := service.GetKeys(pathToInputLines)
    fmt.Printf("Keys: %d\n", len(keys))
    
    unusedKeys := service.GetUnusedKeys(pathForSearching, keys)
    fmt.Printf("Unused keys: %d \n", len(unusedKeys))
}
