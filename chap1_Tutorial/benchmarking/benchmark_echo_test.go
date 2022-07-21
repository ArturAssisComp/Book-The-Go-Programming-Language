package main

import (
    "testing"
    "os/exec"
)


func BenchmarkEfficientEcho(b *testing.B) {
    for n := 0; n < b.N; n++ {
        cmd := exec.Command("./efficient_echo", "1", "2", "\"Hello, World!\"", "abc12")
        cmd.Run()
    }
}

func BenchmarkInefficientEcho1(b *testing.B) {
    for n := 0; n < b.N; n++ {
        cmd := exec.Command("./inefficient_echo1", "1", "2", "\"Hello, World!\"", "abc12")
        cmd.Run()
    }
}

func BenchmarkInefficientEcho2(b *testing.B) {
    for n := 0; n < b.N; n++ {
        cmd := exec.Command("./inefficient_echo1", "1", "2", "\"Hello, World!\"", "abc12")
        cmd.Run()
    }
}
