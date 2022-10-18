package main

import(
	"fmt"
	"time"
	_ "runtime"
)


func main(){

	/*
	go func(){
		defer fmt.Println("A.defer")

		func(){
			defer fmt.Println("B.defer")
			runtime.Goexit() //直接退出func()和外部的func()
			fmt.Println("B")
		}() //定义了这么一个函数 并且直接调用

		fmt.Println("A")
	}()
	*/

	go func(a int,b int) bool{
		fmt.Println("a: ",a)
		fmt.Println("b: ",b)
		return true
	}(10,20)

	//死循环 为了不让线程结束
	for {
		time.Sleep(1*time.Second)
	}

}