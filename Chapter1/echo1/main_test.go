package main

import "testing"

func BenchmarkEcho(b *testing.B) {
    for i := 0; i < b.N; i++ {
        echo()
    }
}
