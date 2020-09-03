package go_pool

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestName1(t *testing.T) {
	//初始化协程池
	pool := Init(10)
	err := pool.Add(nil, F1, "小王")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	time.Sleep(time.Second)
}

func TestName2(t *testing.T) {
	//初始化协程池
	pool := Init(10)
	err := pool.Add(nil, F2, "小张", 18)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	time.Sleep(time.Second)
}

func TestName3(t *testing.T) {
	//初始化协程池
	pool := Init(10)

	wg := sync.WaitGroup{}
	wg.Add(1)

	u := user{}
	u.age = 19
	u.name = "小明"
	err := pool.Add(&wg, F3, u)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	wg.Wait()
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
