package main

import (
	"fmt"
	"runtime"
	"time"
)

func s() {
	runtime.GOMAXPROCS(8)
	n := 10
	defer func(pass int) {
		end := 100
		fmt.Println("run cost", end-n, end-pass)
	}(n)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second)
	n = 1
}
