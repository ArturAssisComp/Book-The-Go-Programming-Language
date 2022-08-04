/*
Name of the program: print_args
Topics covered:
    - Variable short declaration;
    - String concatenation;
    - Loops with range;
    - Importing multiple libs;
    - Basic slicing;
    - Int to string conversion;
*/

//Prints the arguments from CLI.
package main 

import (
    "fmt"
    "os"
    "strconv"
)


func main() {
    fmt.Println("    Command: " + os.Args[0]);
    for index, argument := range os.Args[1:] { //:= is used for short declaration.
        fmt.Println("Argument[" + strconv.Itoa(index + 1) + "]: " + argument)
    }

}
