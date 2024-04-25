package main

import "fmt"

type curImage struct {
}

func (*curImage) showImage() {
	fmt.Print("Image is showing")
}
