package main

type Credit struct {
}

func (*Credit) getNameCard() string {
	return "This is Credit Card"
}
