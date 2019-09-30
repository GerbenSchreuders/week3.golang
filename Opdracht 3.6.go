package main

import "fmt"

var controletekst string
var programmeur1 string
var programmeur2 string
var programmeur3 string
var programmeur4 string
var programmeur5 string

func main(){
	controletekst = "Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk."
	programmeur1 = "Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk."
	programmeur2 = "Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Prorgammeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk."
	programmeur3 = "Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk."
	programmeur4 = "Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk. Programmeren is best leuk."

	if controletekst == programmeur1 {
		fmt.Println("Programmeur 1 = geslaagd")

	} else {
		fmt.Println("Programmeur 1 = gezakt")
	}

	if controletekst == programmeur2 {
		fmt.Println("Programmeur 2 = geslaagd")

	} else {
		fmt.Println("Programmeur 2 = gezakt")
	}

	if controletekst == programmeur3 {
		fmt.Println("Programmeur 3 = geslaagd")

	} else {
		fmt.Println("Programmeur 3 = gezakt")
	}

	if controletekst == programmeur4 {
		fmt.Println("Programmeur 4 = geslaagd")

	} else {
		fmt.Println("Programmeur 4 = gezakt")
	}

	if controletekst == programmeur5 {
		fmt.Println("Programmeur 5 = geslaagd")

	} else {
		fmt.Println("Programmeur 5 = gezakt")
	}


}
