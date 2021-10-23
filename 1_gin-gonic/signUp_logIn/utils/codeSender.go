package utils

import (
	"fmt"
	"net/smtp"
)

func SendCode(receiver, password string) {

	Sender := "YOUR EMAIL"
	Password := "YOUR EMIAL PASSWORD"

	Receivers := []string{
		receiver,
	}

	Host := "smtp.gmail.com"
	Address := "smtp.gmail.com:587"
	//Port := "587"
	fmt.Println(Receivers)
	Auth := smtp.PlainAuth("", Sender, Password, Host)

	Message := []byte(password)

	fmt.Println("Sending ...")

	err := smtp.SendMail(Address, Auth, Sender, Receivers, Message)

	if err != nil {

		panic(err)
	}

	fmt.Println("sended")
	fmt.Println(password)

}
