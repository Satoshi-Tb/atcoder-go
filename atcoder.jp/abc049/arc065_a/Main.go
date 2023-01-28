package main

import (
	"fmt"
	"strings"
)

func main() {
    abc_049_an()
}

func abc_049_an() {
	var s string
    fmt.Scan(&s)


    words := []string {reverse("dreamer"), reverse("dream"), reverse("eraser"), reverse("erase")}
    s = reverse(s)


    outer_loop:
    for {

        for _, w := range words {
            idx := strings.Index(s, w)

            if idx == 0 {
                s = s[len(w):]

                if s == "" {
                    fmt.Println("YES")
                    return
                }
                continue outer_loop
            }
        }

        fmt.Println("NO")
        return
    }
}

func reverse(s string) string {
    rs := []rune(s)
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        rs[i], rs[j] = rs[j], rs[i]
    }
    return string(rs)
}