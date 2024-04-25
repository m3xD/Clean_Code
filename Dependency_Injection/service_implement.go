package main

import "fmt"

type imgFromHN struct {
}

type imgFromHL struct {
}

func (*imgFromHN) GetImage() {
	fmt.Print("Getting image from Hanoi")
}

func (*imgFromHL) GetImage() {
	fmt.Print("Getting image from Halong")
}
