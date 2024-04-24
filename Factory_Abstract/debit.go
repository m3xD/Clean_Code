package main

type Debit struct {
}

func (*Debit) getNameCard() string {
	return "This is Debit Card"
}
