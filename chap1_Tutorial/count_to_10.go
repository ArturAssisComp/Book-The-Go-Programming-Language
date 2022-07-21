/*
Name of the program: count_to_10 
Topics covered:
    - While loop;
    - Traditional for loop;
    - Infinite loop;
    - Variable declaration and initialization;
*/

package main

import (
    "fmt"
    "strconv"
)


func main() {
    var i int
    var max int = 10;

    /* Count to max using different loop styles */
    //Traditional for loop:
    fmt.Println("Tradtional for loop:")
    for i := 0; i < max; i++ {
        fmt.Println(strconv.Itoa(i + 1))
    }


    //while-like loop:
    fmt.Println("While-like for loop:")
    i = 0
    for i < max {
        fmt.Println(strconv.Itoa(i + 1))
        i++;
    }

    //Infinite loop:
    fmt.Println("Infinite loop:")
    i = 0
    for {
        if i >= max {
            break
        }
        fmt.Println(strconv.Itoa(i + 1))
        i++;
    }


}
