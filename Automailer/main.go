package main

import (
	"fmt"

	gomail "gopkg.in/gomail.v2"
)

func main() {
	// Variablen voor de 2 parameters
	var Bestandlocatie string
	var Verstuur string

	fmt.Println("Naar welke email wilt u versturen? ")
	fmt.Scanln(&Verstuur)

	fmt.Println("Welke bestand wilt u sturen? voorbeeld: /home/user/bestand")
	fmt.Scanln(&Bestandlocatie)

	// Variable voor de gegevens
	Sender := "huytran1163@gmail.com"
	Wachtwoord := "skservyxkjekicgy"

	m := gomail.NewMessage()

	// Inhoud van het email
	m.SetHeader("From", Sender)
	m.SetHeader("To", Verstuur)
	m.SetHeader("Subject", "Dit is de automailer")
	m.SetBody("text/plain", "Hier is de bijlage")

	// Verstuurd een email met een bestand.
	m.Attach(Bestandlocatie)

	// Functie NewDialer wordt gebruikt om op de smtp server te connecten en een email te versturen.
	d := gomail.NewDialer("smtp.gmail.com", 587, Sender, Wachtwoord)
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Email verstuurd")
}
