package main

import "fmt"

type singleton struct {
}

var x *singleton

func getInstance() *singleton {
	if x == nil {
		fmt.Print("Init variable now")
		x = &singleton{}
	} else {
		fmt.Println("Variable already created")
	}
	return x
}
