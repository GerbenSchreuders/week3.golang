package main

import (
	"bufio"
	"fmt"
	"os"
)

//Commentaar


func main (){
	USER := bufio.NewReader(os.Stdin)
	fmt.Print("Gebruikersnaam: ")
	RS, _ := USER.ReadString('\n')

	WW := bufio.NewReader(os.Stdin)
	fmt.Print("Wachtwoord: ")
	WS, _ := WW.ReadString('\n')

	if RS == "Bezoeker1" {
		if WS == "12345" {
			fmt.Print("Correct wachtwoord")
		} else {
			fmt.Print("Incorrect wachtwoord, probeer opnieuw")
		}

	} else if RS == "Bezoeker2" {
		if WS == "asdfg" {
			fmt.Print("Correct wachtwoord")
		} else {
			fmt.Print("Incorrect wachtwoord, probeer opnieuw")
		}
	} else if RS == "Bezoeker 3" {
		if WS == "hjkl" {
			fmt.Print("Correct wachtwoord")
		} else {
			fmt.Print("Incorrect wachtwoord, probeer opnieuw")
		}
	} else {
		fmt.Print("Gebruikersnaam onjuist. probeer opnieuw.")
	}

}


