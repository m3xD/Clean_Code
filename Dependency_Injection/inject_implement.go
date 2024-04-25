package main

type getImgFromHN struct {
}

type getImgFromHL struct {
}

func (*getImgFromHN) getIns() Image {
	return &imgFromHN{}
}

func (*getImgFromHL) getIns() Image {
	return &imgFromHL{}
}
