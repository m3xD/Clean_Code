package main

type Card interface {
	getNameCard() string
}

func GetCardFactory(typeOfCard string) Card {
	if typeOfCard == "Credit" {
		return &Credit{}
	} else {
		return &Debit{}
	}
}
