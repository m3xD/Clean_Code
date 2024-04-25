package main

type proxy struct {
}

var img *curImage

func (*proxy) showImage() {
	if img == nil {
		img = &curImage{}
	}
	img.showImage()
}
