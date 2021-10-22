package main

import "fmt"

func main() {
	//var (
	//	strArg string
	//	intArg int
	//	boolArg bool
	//)
	//fmt.Scan(&strArg, &intArg, &boolArg)
	//fmt.Printf("Scan input strArg:%s intArg:%d boolArg:%t \n", strArg, intArg, boolArg)

//	var (
//		strArg string
//		intArg int
//		boolArg bool
//	)
//	fmt.Scanf("strArg:%s intArg:%d boolArg:%t", &strArg, &intArg, &boolArg)
//	fmt.Printf("Scan input strArg:%s intArg:%d boolArg:%t \n", strArg, intArg, boolArg)

	var (
		strArg string
		intArg int
		boolArg bool
	)
	fmt.Scanln(&strArg, &intArg, &boolArg)
	fmt.Printf("Scan input strArg:%s intArg:%d boolArg:%t \n", strArg, intArg, boolArg)
}
