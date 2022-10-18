package main

import(
	"fmt"
	"time"
)

func newTask(){
	i:=0
	for{
		i++
		fmt.Printf("new Goroutine: i = %d\n",i)
		time.Sleep(1*time.Second)
	}
}

func main(){
	//创建一个goroutine 去执行newTask()流程
	go newTask()

	i := 0

	/*
		如果主线程结束，则其他子进程都会结束。
	*/
	for{
		i++
		fmt.Printf("main Goroutine: i = %d\n",i)
		time.Sleep(1*time.Second)
	}
}