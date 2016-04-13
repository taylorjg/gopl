// Exercise1.1 prints its command-line arguments including os.Args[0].
package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    fmt.Printf("Name of the command: %s\n", os.Args[0])
    fmt.Println(strings.Join(os.Args[1:], " "))
}
