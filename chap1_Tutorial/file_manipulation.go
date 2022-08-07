
/*
Topics covered:
- file manipulation
*/

package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "strings"
)

func main() {
    file := os.Args[1]
    //Read the file, loading it into the memory
    data, err := ioutil.ReadFile(file)
    if err != nil {
        fmt.Fprintf(os.Stderr, "file_manipulation: %v\n", err)
    }
    fmt.Println("File content:")
    for _, line := range strings.Split(string(data), "\n") {
        fmt.Println(line)
    }
}
