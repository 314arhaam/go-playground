package main

import (
	"time"
	"fmt"
)

func fibo(num uint64) uint64{
	if num == 0 || num == 1 {
		return 1
	} else {
		return fibo(num - 1) + fibo(num - 2)
	}
}

func main(){
	start_timestamp := float64(time.Now().Unix()) + float64(time.Now().Nanosecond())/1e9
	fmt.Println(fibo(39))
	end_timestamp := float64(time.Now().Unix()) + float64(time.Now().Nanosecond())/1e9
	fmt.Println(end_timestamp - start_timestamp)
}