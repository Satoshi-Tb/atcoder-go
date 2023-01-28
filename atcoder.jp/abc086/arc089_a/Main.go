package main

import (
	"fmt"
)

func main() {
	abc_086_2()
}

type Destination struct {
	time int
	x    int
	y    int
}

func abc_086_2() {
	var n int
	var t, x, y int
	var destList []Destination
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&t, &x, &y)
		var dest Destination
		dest.time = t
		dest.x = x
		dest.y = y
		destList = append(destList, dest)
	}

	var sx, sy, befTime int = 0, 0, 0
	for _, d := range destList {
		var times int = d.time - befTime
		var destx int = abs(d.x - sx)
		var desty int = abs(d.y - sy)

		if times < destx+desty || times%2 != (destx+desty)%2 {
			fmt.Println("No")
			return
		}
		sx = d.x
		sy = d.y
		befTime = d.time
	}

	fmt.Println("Yes")
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	} else {
		return a
	}
}