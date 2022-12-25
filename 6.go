package main

import "fmt"

type x struct {
	time   int
	number int
	z      [3 * 10000]int
}

func main() {
	var a int
	var g [1000000]int

	fmt.Scan(&a)
	var x [10000]x
	for b := 0; b < a; b++ {
		fmt.Scan(&x[b].time)
		fmt.Scan(&x[b].number)
		for c := 0; c < x[b].number; c++ {
			fmt.Scan(&x[b].z[c])
		}
	}

	for b := 0; b < a; b++ {
		var gg = 0
		for c := 0; c < a; c++ {
			if x[b].time-x[c].time < 86400 && x[b].time-x[c].time >= 0 {
				for d := 0; d < x[c].number; d++ {
					g[gg] = x[c].z[d]
					gg++
				}
			}
		}
		var w []int
		w = append(w, g[0])
		for d := 0; d < gg; d++ {
			j := true
			for c := 0; c < len(w); c++ {
				if g[d] == w[c] {
					j = false
				}
			}
			if j == true {
				w = append(w, g[d])
			}
		}
		fmt.Println(len(w))
	}

}
