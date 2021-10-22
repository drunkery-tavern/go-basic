package main

import "fmt"

var (
	floatValue  = 3.1415926
	intValue    = 501
	stringValue = "Go fmt"
	boolValue   = false
)

type Player struct {
	name string
	age  int
}

func main() {
	//fmt.Printf("intValue printf with %%v：%v\n", intValue)   // intValue printf with %v：501
	//fmt.Printf("boolValue printf with %%v：%v\n", boolValue) //boolValue printf with %v：false
	//
	//player := Player{name: "drunkery", age: 23}
	//fmt.Printf("player struct printf with %%v：%v\n", player)   //player struct printf with %v：{drunkery 23}
	//fmt.Printf("player struct printf with %%#v：%#v\n", player) //player struct printf with %#v：main.Player{name:"drunkery", age:23}
	//fmt.Printf("player struct printf with %%T：%T\n", player)   //player struct printf with %T：main.Player
	//
	//fmt.Printf("boolValue printf with %%t：%t\n", boolValue) //boolValue printf with %t：false

	//fmt.Printf("intValue printf with %%b:%b\n", intValue) //intValue printf with %b:111110101
	//fmt.Printf("intValue printf with %%c:%c\n", intValue) //intValue printf with %c:ǵ
	//fmt.Printf("intValue printf with %%d:%d\n", intValue) //intValue printf with %d:501
	//fmt.Printf("intValue printf with %%o:%o\n", intValue) //intValue printf with %o:765
	//fmt.Printf("intValue printf with %%x:%x\n", intValue) //intValue printf with %x:1f5
	//fmt.Printf("intValue printf with %%X:%X\n", intValue) //intValue printf with %X:1F5

	//fmt.Printf("floatValue printf with %%b:%b\n", floatValue) //floatValue printf with %b:7074237631354954p-51
	//fmt.Printf("floatValue printf with %%e:%e\n", floatValue) //floatValue printf with %e:3.141593e+00
	//fmt.Printf("floatValue printf with %%E:%E\n", floatValue) //floatValue printf with %E:3.141593E+00
	//fmt.Printf("floatValue printf with %%f:%f\n", floatValue) //floatValue printf with %f:3.141593
	//fmt.Printf("floatValue printf with %%g:%g\n", floatValue) //floatValue printf with %g:3.1415926
	//fmt.Printf("floatValue printf with %%G:%G\n", floatValue) //floatValue printf with %G:3.1415926

	//fmt.Printf("stringValue printf with %%s:%s\n", stringValue) //stringValue printf with %s:The format 'verbs' are derived from C's but are simpler.
	//fmt.Printf("stringValue printf with %%q:%q\n", stringValue) //stringValue printf with %q:"The format 'verbs' are derived from C's but are simpler."
	//fmt.Printf("stringValue printf with %%x:%x\n", stringValue) //stringValue printf with %x:54686520666f726d617420277665726273272061726520646572697665642066726f6d2043277320627574206172652073696d706c65722e
	//fmt.Printf("stringValue printf with %%X:%X\n", stringValue) //stringValue printf with %X:54686520666F726D617420277665726273272061726520646572697665642066726F6D2043277320627574206172652073696D706C65722E

	//fmt.Printf("intValue printf with %%p：%p\n", &intValue)   //intValue printf with %p：0x164238
	//fmt.Printf("intValue printf with %%#p：%#p\n", &intValue)   //intValue printf with %#p：164238

	fmt.Printf("floatValue printf with %%f:%f\n", floatValue) //floatValue printf with %f:3.141593
	fmt.Printf("floatValue printf with %%9f:%9f\n", floatValue) //floatValue printf with %9f: 3.141593
	fmt.Printf("floatValue printf with %%.2f:%.2f\n", floatValue) //floatValue printf with %.2f:3.14
	fmt.Printf("floatValue printf with %%9.2f:%9.2f\n", floatValue) //floatValue printf with %9.2f:     3.14
	fmt.Printf("floatValue printf with %%9.f:%9.f\n", floatValue) //floatValue printf with %9.f:        3

}
