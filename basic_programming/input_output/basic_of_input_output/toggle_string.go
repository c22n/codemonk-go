package main

import "fmt"
import "strings"

func main() {
    swapCase := func(r rune) rune {
        switch {
        case r >= 'A' && r <= 'Z':
            return r + 32
        case r >= 'a' && r <= 'z':
            return r - 32
        }
        return r
    }
    var S string
    fmt.Scan(&S)
    fmt.Println(strings.Map(swapCase, S))
}
