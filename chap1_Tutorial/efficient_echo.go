/*
Topics covered:
    - Joining strings;
*/

package main

import (
    "fmt"
    "strings"
    "os"
)

func main() {
    fmt.Println(strings.Join(os.Args[1:], " "))
}
