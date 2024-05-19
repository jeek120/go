package main

import (
	"fmt"
	"time"
)

type a struct {
	n int
}

func main() {
	c := make(chan a, 2)
	go func() {
		b := a{n: 1}
		d := a{n: 2}
		e := a{n: 3}
		f := a{n: 4}
		g := a{n: 5}
		c <- b
		go func() {
			c <- d
			c <- e
			fmt.Println("携程1完成")
		}()
		go func() {
			c <- f
			c <- g
			fmt.Println("携程2完成")
		}()
		time.Sleep(2 * time.Second)
		//close(c)
		fmt.Println("chan关闭")
		fmt.Printf("%v\n", <-c)
		fmt.Printf("%v\n", <-c)
		fmt.Printf("%v\n", <-c)
		fmt.Printf("%v\n", <-c)
		fmt.Printf("%v\n", <-c)
		fmt.Printf("%v\n", <-c)
		fmt.Printf("%v\n", <-c)
	}()
	time.Sleep(500 * time.Second)
}
