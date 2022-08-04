/*
Topics covered:
- map;
- if/else conditionals
- passing argument by reference
- error checking
- file manipulation
- Printf
- Declaring and defining function
- range through map items
*/

// This program prints the count and the text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    /*
    => map: creates a new map (like a dictionary) using the built-in function 'make'.
    A map is data structure that stores a value for a given key. The key may be
    anything that can be compared using ==. The value may be anything. Its store,
    retrieve, and test for an item in the set are done in constant time.
    */
    counts := make(map[string]int)
    files := os.Args[1:]

    //=>if/else conditionals
    if len(files) == 0 {
        //=> passing argument by reference: counts is passed by reference.
        countLines(os.Stdin, counts) 
    } else{
        for _, arg := range files {
            //=>file manipulation: file must be opened before use and closed
            //after use.
            f, err := os.Open(arg)
            if err != nil { //=> Error checking;
                fmt.Fprintf(os.Stderr, "dup_2: %v\n", err)
                continue //=> continue keyword;
            }
            countLines(f, counts)
            f.Close()
        }
    }
    //=> range through map items: the items are unpacked in random order from the
    //map.
    for line, n := range counts {
        if n > 1 {
            //=>Printf: C-like print function. It allows formatting string.
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
//=> Declaring and defining function: A function can be declared and defined
//after a line in which it is called.
func countLines(f *os.File, counts map[string]int) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        counts[input.Text()]++
    }
    //OBS: input.Err() should be checked for errors.
}
