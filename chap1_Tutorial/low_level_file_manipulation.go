
/*
Topics covered:
- error checking
- low level file manipulation
- input scan
- Fprintf
*/

package main

import (
    "fmt"
    "os"
    "bufio"
)

func main() {
    file := os.Args[1]
    f, err := os.Open(file)
    if err != nil {
        fmt.Fprintf(os.Stderr, "low_level_file_manipulation: %v\n", err)
    }
    input := bufio.NewScanner(f)
    fmt.Println("File content: ")
    for input.Scan() {
        fmt.Println(input.Text())
    }
    //OBS: errors from input.Err() should be checked.
    f.Close()
}
