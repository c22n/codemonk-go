package main

import "fmt"
import "os"

func main() {
    var x string
    fmt.Scan(&x)
    
    for i, j := 0, len([]rune(x))-1; j - i > 0; i, j = i+1, j-1 {
        if x[i] != x[j] {
            fmt.Println("NO")
            os.Exit(0)
        }
    }
    fmt.Println("YES")
}
