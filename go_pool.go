package go_pool

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

/*
	协程池
*/

type Work struct {
	usable   int64 //可用数量
	totalNum int64 //总数量
	pool     chan obj
}

type obj struct {
	Fun func([]reflect.Value) []reflect.Value
	val []reflect.Value
	wg  *sync.WaitGroup
}

/*
	初始化池
*/
func Init(num int64) *Work {
	work := &Work{}
	work.pool = make(chan obj, num)
	work.totalNum = num
	for i := int64(0); i < num; i++ {
		go work.run()
	}
	return work
}

/*
	奔跑
*/
func (this *Work) run() {
	for ; ; {
		select {
		case fun := <-this.pool:
			this.usable++
			fun.callFunc()
			this.usable--
		}
	}
}

/*
	调用方法
*/
func (this obj) callFunc() {
	this.Fun(this.val)
	if this.wg != nil {
		this.wg.Done()
	}
}

/*
	添加异步方法
	支持WaitGroup同步，当不需要WaitGroup，传入nil
*/
func (this *Work) Add(wg *sync.WaitGroup, fun interface{}, val ...interface{}) error {
	//校验参数
	if err := verifyParams(fun, val...); err != nil {
		return err
	}

	//获取参数
	params, err := getParams(fun, val...)
	if err != nil {
		return err
	}

	//调用方法
	this.runGo(wg, fun, params)
	return nil
}

/*
	校验参数
*/
func verifyParams(fun interface{}, val ...interface{}) error {
	funcType := reflect.TypeOf(fun)
	//校验参数数量
	if funcType.NumIn() > len(val) {
		return errors.New(fmt.Sprintf("参数过少，方法参数数量有 %d 个，您输入的参数数量为 %d 个", funcType.NumIn(), len(val)))
	} else if funcType.NumIn() < len(val) {
		return errors.New(fmt.Sprintf("参数过多，方法参数数量只有 %d 个，您输入的参数数量为 %d 个", funcType.NumIn(), len(val)))
	}
	//校验参数类型
	for i := 0; i < funcType.NumIn(); i++ {
		paramType := reflect.TypeOf(val[i])
		if funcType.In(i) != paramType {
			return errors.New("正确参数类型为:" + funcType.In(i).String() + "，您输入类型为:" + paramType.String())
		}
	}
	return nil
}

/*
	提取参数
*/
func getParams(fun interface{}, val ...interface{}) ([]reflect.Value, error) {
	funcType := reflect.TypeOf(fun)
	params := make([]reflect.Value, funcType.NumIn())
	for i := 0; i < funcType.NumIn(); i++ {
		params[i] = reflect.ValueOf(val[i])
	}
	return params, nil
}

/*
	整理数据并投递到异步处理池进行处理
*/
func (this *Work) runGo(wg *sync.WaitGroup, fun interface{}, value []reflect.Value) {
	funcValue := reflect.ValueOf(fun)
	funObj := obj{}
	if wg != nil {
		funObj.wg = wg
	} else {
		funObj.wg = nil
	}
	funObj.Fun = funcValue.Call
	funObj.val = value
	this.pool <- funObj
}

/*
	可用的协程数量
*/
func (this *Work) GetUsableNum() int64 {
	num := this.usable
	return num
}

/*
	协程总数量
*/
func (this *Work) GetTotalNum() int64 {
	num := this.totalNum
	return num
}
