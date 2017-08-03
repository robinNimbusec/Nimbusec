package main 

import (
    "fmt"
    "../findUnusedProps/services"
    "os"
    "flag"
    "path/filepath"
)

var paths []string

func visit(path string, f os.FileInfo, err error) error {
  if filepath.Ext(path) == ".html"{
    paths = append(paths, path)
  }
  return nil
} 

func main(){
    flag.Parse()
    root := flag.Arg(1)
    filepath.Walk(root, visit)

    pathToInputLines := os.Args[1]

    keys := service.GetKeys(pathToInputLines)

    fmt.Printf("Keys: %d\n", len(keys))
    for i := 0; i < len(paths); i++ { 
        keys = service.GetUnusedKeys(paths[i], keys)
    }
    fmt.Printf("Unused keys: %d \n", len(keys))

    service.WriteSliceToFile(keys)
}

