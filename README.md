# go_pool
 go协程池，简单，容易使用

测试方法

    func F0() {
    	fmt.Println("哈哈")
    }

    func F1(name string) {
    	fmt.Println(name)
    }
    
    func F2(name string, age int) {
    	fmt.Println(name, age)
    }
    
    type user struct {
    	name string
    	age  int
    }
    
    func F3(u user) {
    	fmt.Println(u)
    }

无参数调用
   
    //协程数量
    var num int64=10
    //初始化协程池
    pool := Init(num)
    //调用方法
    err := pool.Add(nil, F0)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
     time.Sleep(time.Second)

内部方法调用

    //协程数量
    var num int64=10
    //初始化协程池
    pool := Init(num)
    inFunc:= func() {
        fmt.Println("内部方法")
    }
    //调用方法
    err := pool.Add(nil, inFunc)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
     time.Sleep(time.Second)

单参数调用
   
    //协程数量
    var num int64=10
    //初始化协程池
    pool := Init(num)
    //调用方法
	err := pool.Add(nil, F1, "小明")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	time.Sleep(time.Second)
	
多参数调用
    
    //协程数量
    var num int64=10
    //初始化协程池
    pool := Init(num)
    //调用方法
	err := pool.Add(nil, F2, "小张", 18)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	time.Sleep(time.Second)
	
结构体参数调用，并使用协程同步

    //协程数量
    var num int64=10
    //初始化协程池
    pool := Init(num)
	//使用协程同步
	wg := sync.WaitGroup{}
	wg.Add(1)

	u := user{}
	u.age = 19
	u.name = "小红"
	err := pool.Add(&wg, F3, u)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	wg.Wait()
	
	
~~~~
作者：tom
联系方式（微信同步）：13229025103
