package main

import "fmt"

var isBoterAanwezig, isKaasAanwezig, isMelkAanwezig, isEiAanwezig bool

func main() {
	isBoterAanwezig = true
	isKaasAanwezig = true
	isEiAanwezig = false
	isMelkAanwezig = false

	fmt.Println("Boter aanwezig\t",isBoterAanwezig, "\nEi aanwezig\t", isEiAanwezig, "Ei aanwezig\t", isEiAanwezig)
	fmt.Println("+TESTS+")
	fmt.Print("test1\t\t")

	if isBoterAanwezig == true && isKaasAanwezig == true {
		fmt.Println("Boter en kaas aanwezig")

	} else {
		fmt.Println("Beide niet aanwezig")
	}

	if isBoterAanwezig == true && isEiAanwezig == true {
		fmt.Println("Boter en ei aanwezig")

	} else {
		fmt.Println("Beide niet aanwezig")
	}

	if isBoterAanwezig == true && isMelkAanwezig == true {
		fmt.Println("Boter en melk aanwezig")

	} else {
		fmt.Println("Beide niet aanwezig")
	}

	if isBoterAanwezig == true && isMelkAanwezig == true {
		fmt.Println("Boter en melk aanwezig")

	} else {
		fmt.Println("Beide niet aanwezig")
	}


}
