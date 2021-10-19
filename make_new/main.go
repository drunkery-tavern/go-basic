package main

import (
	"fmt"
	"strconv"
	"sync"
)

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

/*func main() {
	var ch chan int
	//a := <-ch
	//fmt.Println(a)

	//ch <- 10
	close(ch)
}*/

/*func main() {
	//新建计时器，5秒后触发
	timer := time.NewTimer(5 * time.Second)
	//新开启一个线程来处理触发后的事件
	go func() {
		//等触发时的信号
		<-timer.C
		fmt.Println("Timer 结束。。")
	}()

	//由于上面的等待信号是在新线程中，所以代码会继续往下执行，停掉计时器
	time.Sleep(3*time.Second)
	stop := timer.Stop()
	if stop {
		fmt.Println("Timer 停止。。")
	}
}*/

/*func main() {
	//返回一个通道：chan，存储的是d时间间隔后的当前时间。
	ch := time.After(3 * time.Second) //3s后
	fmt.Printf("%T\n", ch) // <-chan time.Time
	fmt.Println(time.Now()) //2021-10-12 10:03:18.8830048 +0800 CST m=+0.001656201
	t := <-ch //阻塞3s
	fmt.Println(t) //2021-10-12 10:03:21.8917384 +0800 CST m=+3.010389801
}*/

/*func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 200
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 100
	}()

	select {
	case num1 := <-ch1:
		fmt.Println("ch1中取数据。。", num1)
	case num2, ok := <-ch2:
		if ok {
			fmt.Println("ch2中取数据。。", num2)
		} else {
			fmt.Println("ch2通道已经关闭。。")
		}
	}
}
*/

/*func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	//go func() {
	//	ch1 <- 100
	//}()

	select {
	case <-ch1:
		fmt.Println("case1可以执行。。")
	case <-ch2:
		fmt.Println("case2可以执行。。")
	case <-time.After(3 * time.Second):
		fmt.Println("case3执行。。timeout。。")
		default:
			fmt.Println("执行了default。。")
	}
}*/

/*//全局变量
var ticket = 10 // 100张票

func main() {
	go saleTickets("售票口1") // g1,100
	go saleTickets("售票口2") // g2,100
	go saleTickets("售票口3") //g3,100
	go saleTickets("售票口4") //g4,100

	time.Sleep(5*time.Second)
}

func saleTickets(name string) {
	rand.Seed(time.Now().UnixNano())
	//for i:=1;i<=100;i++{
	//	fmt.Println(name,"售出：",i)
	//}
	for { //ticket=1
		if ticket > 0 { //g1,g3,g2,g4
			//睡眠
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			// g1 ,g3, g2,g4
			fmt.Println(name, "售出：", ticket)  // 1 , 0, -1 , -2
			ticket--   //0 , -1 ,-2 , -3
		} else {
			fmt.Println(name,"售罄，没有票了。。")
			break
		}
	}
}*/

/*//全局变量
var ticket = 10 // 100张票

var wg sync.WaitGroup
var matex sync.Mutex // 创建锁头

func main() {
	wg.Add(4)
	go saleTickets("售票口1") // g1,100
	go saleTickets("售票口2") // g2,100
	go saleTickets("售票口3") //g3,100
	go saleTickets("售票口4") //g4,100
	wg.Wait()              // main要等待。。。

	//time.Sleep(5*time.Second)
}

func saleTickets(name string) {
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	//for i:=1;i<=100;i++{
	//	fmt.Println(name,"售出：",i)
	//}
	for { //ticket=1
		matex.Lock()
		if ticket > 0 { //g1,g3,g2,g4
			//睡眠
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			// g1 ,g3, g2,g4
			fmt.Println(name, "售出：", ticket) // 1 , 0, -1 , -2
			ticket--                         //0 , -1 ,-2 , -3
		} else {
			matex.Unlock() //解锁
			fmt.Println(name, "售罄，没有票了。。")
			break
		}
		matex.Unlock() //解锁
	}
}*/

/*var wg sync.WaitGroup // 创建同步等待组对象
func main()  {
	//设置等待组中，要执行的goroutine的数量
	wg.Add(2)
	go fun1()
	go fun2()
	fmt.Println("main进入阻塞状态。。。等待wg中的子goroutine结束。。")
	wg.Wait() //表示main goroutine进入等待，意味着阻塞
	fmt.Println("main，解除阻塞。。")
}
func fun1()  {
	for i:=1;i<=10;i++{
		fmt.Println("fun1.。。i:",i)
	}
	wg.Done() //给wg等待中的执行的goroutine数量减1.同Add(-1)
}
func fun2()  {
	defer wg.Done()
	for j:=1;j<=10;j++{
		fmt.Println("\tfun2..j,",j)
	}
}*/

/*var (
	rwMutex sync.RWMutex
	wg sync.WaitGroup
)

func main() {
	//wg.Add(2)
	//
	//多个同时读取
	//go readData(1)
	//go readData(2)

	wg.Add(3)
	go writeData(1)
	go readData(2)
	go writeData(3)

	wg.Wait()
	fmt.Println("main..over...")
}


func writeData(i int){
	defer wg.Done()
	fmt.Println(i,"开始写：write start。。")
	rwMutex.Lock()//写操作上锁
	fmt.Println(i,"正在写：writing。。。。")
	time.Sleep(3*time.Second)
	rwMutex.Unlock()
	fmt.Println(i,"写结束：write over。。")
}

func readData(i int) {
	defer wg.Done()
	fmt.Println(i, "开始读：read start。。")
	rwMutex.RLock() //读操作上锁
	fmt.Println(i,"正在读取数据：reading。。。")
	time.Sleep(3*time.Second)
	rwMutex.RUnlock() //读操作解锁
	fmt.Println(i,"读结束：read over。。。")
}*/
var m = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
