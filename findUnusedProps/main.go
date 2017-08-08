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

    var messageProps []string
    for i := 0; i < len(paths); i++ { 
        messageProps = append(messageProps, service.GetMsgProps(paths[i])...)
    }


    ukeys := service.GetUnusedKeys(messageProps, keys)

    service.GetMsgProps(paths[0])

    service.WriteSliceToFile(ukeys)
}

