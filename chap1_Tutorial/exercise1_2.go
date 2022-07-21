//Prints the arguments from CLI.
package main 

import (
    "fmt"
    "os"
    "strconv"
)


func main() {
    for index, argument := range os.Args[1:] { //:= is used for short declaration.
        fmt.Println("Argument[" + strconv.Itoa(index + 1) + "]: " + argument)
    }
}
