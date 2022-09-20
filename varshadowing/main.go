package main

import (
	"errors"
	"fmt"
)

var a int = 2020

func checkYear() error {
	err := errors.New("wrong year")

	// switch 控制语句中局部变量 a 遮蔽了外层包代码级变量 a
	// 修改方法：修改作用域
	/*
		a, err := getYear()
		switch a {
	*/
	switch a, err := getYear(); a {
	case 2020:
		fmt.Println("it is", a, err)
	case 2021:
		fmt.Println("it is", a)
		err = nil // switch 控制语句中局部变量 err 遮蔽了外层包代码级变量 err; 这里并没有影响外部 err 的值
	}
	fmt.Println("after check, it is", a)
	return err
}

// 遮蔽了标识符 new
type new int

func getYear() (new, error) {
	var b int16 = 2021
	return new(b), nil
}

func main() {
	err := checkYear()

	if err != nil {
		fmt.Println("call checkYear error:", err)
		return
	}

	fmt.Println("call checkYear ok")
}
