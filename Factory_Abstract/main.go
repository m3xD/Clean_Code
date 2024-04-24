package main

import "fmt"

func main() {
	cardFactoryCredit := GetCardFactory("credit")
	fmt.Print(cardFactoryCredit.getNameCard())
	fmt.Println()
	cardFactoryDebit := GetCardFactory("credit")
	fmt.Print(cardFactoryDebit.getNameCard())
}
