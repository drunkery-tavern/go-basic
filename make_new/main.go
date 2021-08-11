package main

import "log"

func main() {

	i := new(int64)
	log.Println(*i) // 0

	m := new(map[int]int)
	log.Println(*m) //map[]
	(*m)[1] = 1     // panic: assignment to entry in nil map

}
