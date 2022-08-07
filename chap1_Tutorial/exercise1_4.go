
// This program prints the count and the text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
package main

import (
    "bufio"
    "fmt"
    "os"
)

type container struct {
    filenameSet map[string]bool
    n int
}

func main() {
    counts := make(map[string]*container)
    files := os.Args[1:]

    if len(files) == 0 {
        countLines(os.Stdin, counts, "") 
    } else{
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil { 
                fmt.Fprintf(os.Stderr, "exercise1_4: %v\n", err)
                continue 
            }
            countLines(f, counts, arg)
            f.Close()
        }
    }
    for line, struct_object := range counts {
        if struct_object.n > 1 {
        keys := make([]string, len(struct_object.filenameSet))
        i := 0
        for key, _ := range struct_object.filenameSet {
            keys[i] = key
            i++
        }
            fmt.Printf("%d\t%s\tFiles-> %v\n", struct_object.n, line, keys)
        }
    }
}
func countLines(f *os.File, counts map[string]*container, filename string) {
    input := bufio.NewScanner(f)
    for input.Scan() {

        inputText := input.Text()
        inputText = inputText
        if counts[inputText] == nil {
            counts[inputText] = &container{map[string]bool{}, 0}
        }
        counts[inputText].n++
        counts[inputText].filenameSet[filename] = true
    }
    //OBS: input.Err() should be checked for errors.
}
