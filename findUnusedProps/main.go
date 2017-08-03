package main 

import (
    "fmt"
    "../findUnusedProps/services"
    "os"
)

func main() {
    pathToInputLines := os.Args[1]
    pathForSearching := os.Args[2]

    keys := service.GetKeys(pathToInputLines)
    fmt.Printf("Keys: %d\n", len(keys))
    
    unusedKeys := service.GetUnusedKeys(pathForSearching, keys)
    fmt.Printf("Unused keys: %d \n", len(unusedKeys))

    service.WriteSliceToFile(unusedKeys)
}
