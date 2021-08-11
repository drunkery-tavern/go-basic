package main

import "log"

func main() {

	i := new(int64)
	log.Println(*i) // 0

	m := new(map[int]int)
	log.Println(*m)        //map[]
	log.Println(*m == nil) //true
	//(*m)[1] = 1     // panic: assignment to entry in nil map

	slice := new([]int)
	log.Println(*slice)        //[]
	log.Println(*slice == nil) //true

	m2 := make(map[int]int, 3)
	log.Println(m2)        //map[]
	log.Println(m2 == nil) //false
	m2[1] = 1
	log.Println(m2[1]) //1
}
