package main

import (
	"fmt"
)

func main(){
	c:=make(chan int)

	go func(){
		for i:=0;i<5;i++{
			c<-i
		}
		close(c)
	}()

	/*
	for {
		//ok如果为true则没有关闭 如果为false则关闭了
		if data,ok := <-c; ok{
			fmt.Println(data)
		}else{
			break
		}
	}
	*/

	//可以用range来迭代不断操作的channel
	for data := range c{
		fmt.Println(data)
	}

	fmt.Println("Main Finished")

}