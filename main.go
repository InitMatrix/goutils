package main

import (
	"fmt"
	"time"
)

/**
 *@Author: yepeng
 *@Date: 2022/11/3 23:50
 *@Name: haha
 *@Desc:
 */
func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("i ----> ", i)
		time.Sleep(1 * time.Second)
	}
}

