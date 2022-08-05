/*
Topics included:
- map;
- access not assigned key value
- range through map items;
- Printf
*/

package main

import (
    "fmt"
)

func main() {
    //=>map: a map with string:string items is declared.
    target_map := make(map[string]string)
    var i int

    //Assign values to some keys:
    target_map["k1"] = "v1"
    target_map["k2"] = "v2"
    target_map["k3"] = "v3"
    target_map["k4"] = "v4"


    //Print the assigned values:
    i = 1
    for key, value := range target_map {
        fmt.Printf("(%d) key: %s , value: %s\n", i, key, value)
        i++
    }

    //Try to print a value that was not assigned. The result will be empty string.
    fmt.Printf("Not yet assigned value for key \"hello\": '%s'\n", target_map["hello"])
}
