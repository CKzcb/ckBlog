/**
 * @Author: 蛋白质先生
 * @Description:
 * @File: test_runtime
 * @Version: 1.0.0
 * @Date: 2023/3/16 08:05
 */

package main

import (
	"fmt"
	"runtime"
	"time"
)

func f1() {
	fmt.Println("in f1 >> ")
	a := make([]uintptr, 2)
	fmt.Println(runtime.Callers(0, a))

	b := runtime.CallersFrames(a[:2])
	c, _ := b.Next()
	fmt.Println(c.Line, c.Function, c.File)
	fmt.Println("out f1 << ", a, time.Now().Local().UnixNano())
}

func f2() {
	f1()
}

func main() {
	f2()
}
