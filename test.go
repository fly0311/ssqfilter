package main

import "fmt"

func main() {
	rs := make(map[int]bool)
	d := make([]int, 33)
	for i := 0; i < 6; i++ {
		d[i] = 1
	}

	for {
		cg := false
		for j := 0; j < 32; j++ {
			if d[j] == 1 && d[j+1] == 0 {
				d[j], d[j+1] = d[j+1], d[j]

				s := 0
				e := j - 1

				ones := 0
				for a := s; a <= e; a++ {
					if d[a] == 1 {
						ones++
					}
				}

				for b := 0; b < ones; b++ {
					d[b] = 1
				}
				for c := ones; c <= e; c++ {
					d[c] = 0
				}

				var this int = 0
				for index := 0; index < 33; index++ {
					if d[index] == 1 {
						this += 1 << (33 - index - 1)
					}
				}
				rs[this] = true
				cg = true
				break
			}
		}
		if cg == false {
			break
		}
	}

	fmt.Println(len(rs))

}
