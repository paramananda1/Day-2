package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func eMailInGolang() {
	fmt.Println("Do you want to send a Email (y/n):")
	var sendemail string
	fmt.Scanf("%s",&sendemail)

	if sendemail == "Y" || sendemail == "y"{

			var to,from,subject,body string
			/*
			fmt.Println("Enter Body to email . Or Enter to send a empty email")
			fmt.Scanf("%s",&body)
			fmt.Println("Enter Subject to email . Or Enter for defult")
			fmt.Scanf("%s",&subject)
			fmt.Println("Enter To Address")
			fmt.Scanf("%s",&to)
			fmt.Println("Enter From Address")
			fmt.Scanf("%s",&from)
			send(to,from,subject,body)
			*/
		to = "paramanandaexample@gamil.com"
		from = "paramanandaexample@gamil.com"
		subject = "test email"
		body= "test example"
		send(to,from,subject,body)

	}
}
func send(to,from,subject,body string) {

	if len(to)<1 || len(from)<1 {
		fmt.Println("Unable to send Email as Either TO of From address wrong. Currect and try again.")
		return
	}
	if len(subject)<1{
		subject="Hello there\n\n"
	}
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject +"\r\n" +
		body +"\r\n"

	auth := smtp.PlainAuth("", "paramanandaexample@gamil.com", "paramananda", "smtp.gmail.com")
	err := smtp.SendMail("smtp.gmail.com:587",
		auth,
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent, visit http://foobarbazz.mailinator.com")
}
