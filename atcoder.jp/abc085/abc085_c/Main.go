package main

import (
	"fmt"
)

func main() {
    abc_085()
}

func abc_085() {
    var n, y int
    fmt.Scan(&n, &y)


    for i := 0; i <= n; i++ {
        if i > n || i * 10000 > y {
            break
        }
        for j := 0; j <= n; j++ {
            if (i + j) > n || (i * 10000 + j * 5000) > y {
                break
            }
            k := n - i - j
            if i * 10000 + j * 5000 + k * 1000 == y && (i + j + k) == n {
                fmt.Printf("%d %d %d\n", i, j, k)
                return
            }
        }
    }

    fmt.Println("-1 -1 -1")
}