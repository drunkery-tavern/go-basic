package main

/*func main() {

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
	log.Println(m2) //map[1:1]
	modifyMap(m2)
	log.Println(m2) //map[1:2]

	channel := make(chan string, 1)
	log.Println(channel)        //0xc00004e1e0
	log.Println(channel == nil) //false
	channel <- "string"
	log.Println(<-channel) //string
	modifyChannel(channel)
	log.Println(<-channel) //new string

}

func modifyMap(m map[int]int) {
	m[1] = 2
}

func modifyChannel(c chan string) {
	c <- "new string"
}*/

/*type student struct {
	name string
	age  int
}

func (s student) newStudent() {

}

func main() {
	var s *student
	s.newStudent()
	m := make(map[string]*student)
	stus := []student{
		{name: "张三", age: 18},
		{name: "李四", age: 23},
		{name: "王五", age: 25},
	}

	for _, stu := range stus {
		fmt.Printf("stu:%v \n", &stu)
		m[stu.name] = &stu
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}*/

/*func main() {
	// *int类型空指针
	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil()) //var a *int IsNil: true
	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid()) //nil IsValid: false
	// 实例化一个匿名结构体
	b := struct{}{}
	// 尝试从结构体中查找"abc"字段
	fmt.Println("b结构体是否存在成员abc:", reflect.ValueOf(b).FieldByName("abc").IsValid()) //b结构体是否存在成员abc: false
	// 尝试从结构体中查找"abc"方法
	fmt.Println("b结构体是否存在方法abc:", reflect.ValueOf(b).MethodByName("abc").IsValid()) //b结构体是否存在方法abc: false
	// map
	c := map[string]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map中是否存在键张三：", reflect.ValueOf(c).MapIndex(reflect.ValueOf("张三")).IsValid()) //map中是否存在键张三： false
}*/

/*type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

// 给student添加两个方法 Study和Sleep(注意首字母大写)
func (s student) Study() string {
	msg := "学习"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "睡觉"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		args := []reflect.Value{}
		v.Method(i).Call(args)
	}
}

func main() {
	stu := student{
		Name:  "张三",
		Score: 90,
	}

	t := reflect.TypeOf(stu)
	fmt.Println(t.Name(), t.Kind()) // student struct
	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}

	printMethod(stu)
}*/

/*func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("goroutine。。。")
		}

	}()

	for i := 0; i < 4; i++ {
		//让出时间片，先让别的协议执行，它执行完，再回来执行此协程
		runtime.Gosched()
		fmt.Println("main。。")
	}
}*/

/*func main() {
	//创建新建的协程
	go func() {
		fmt.Println("goroutine开始。。。")

		//调用了别的函数
		fun()

		fmt.Println("goroutine结束。。")
	}() //别忘了()

	//睡一会儿，不让主协程结束
	time.Sleep(3*time.Second)
}*/
/*func main() {
	var ch1 chan bool       //声明，没有创建
	fmt.Println(ch1)        //<nil>
	fmt.Printf("%T\n", ch1) //chan bool
	ch1 = make(chan bool)   //0xc0000a4000,是引用类型的数据
	fmt.Println(ch1)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("子goroutine中，i：", i)
		}
		// 循环结束后，向通道中写数据，表示要结束了。。
		ch1 <- true

		fmt.Println("结束。。")

	}()

	data := <-ch1 // 从ch1通道中读取数据
	fmt.Println("data-->", data)
	fmt.Println("main。。over。。。。")
}*/

/*func fun() {
	defer fmt.Println("defer。。。")

	//return           //终止此函数
	runtime.Goexit() //终止所在的协程

	fmt.Println("fun函数。。。")
}*/

func main() {
	var ch chan int
	//a := <-ch
	//fmt.Println(a)

	//ch <- 10
	close(ch)


}
