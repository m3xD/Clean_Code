package main

func main() {
	getImgHN := getImgFromHN{}
	getImgHL := getImgFromHL{}

	img1 := getImgHN.getIns()
	img2 := getImgHL.getIns()

	img1.GetImage()
	img2.GetImage()
}
