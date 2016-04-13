// Exercise1.2 prints its command-line arguments with indices.
package main

import (
    "fmt"
    "os"
)

func main() {
    for idx, arg := range os.Args[1:] {
        fmt.Printf("os.Args[%d]: %s\n", idx + 1, arg)
    }
}
