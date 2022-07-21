/*
Topics covered:
    - Short declaration and initialization for multiple variables (like python unpacking)
    - Variable _ to store unused value;
*/

package main

import (
    "fmt"
    "os"
)

func main() {
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
}
