package smtp

import (
	"fmt"
	"net"
	"os"
)

const (
	smtpAddr string = "0.0.0.0:2255"
	greeting string = "220 smtp.testserver.com ESMTP postfix"
)

type smtpSession struct {
	senderAddress          string
	recipientAddress       string
	data                   string
	clientReportedHostname string
	sourceAddress          string
}

//StartSMTPServer starts the main connection loop
func StartSMTPServer() {
	//Start message
	//setup connection listener
	listener, err := net.Listen("tcp", smtpAddr)
	fmt.Println("SMTP server started...")
	errorHandler(err)

	//Main listen loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Something went wrong with the SMTP listener")
		}
		fmt.Println("Incoming...")
		smtpSessionHandler(conn)
	}
}

func smtpSessionHandler(conn net.Conn) {
	session := smtpSession{}
	session.sourceAddress = conn.RemoteAddr().String()
	conn.Write([]byte(greeting))
	//Get the input from the connection

	//Send HELO

	//Parse the first word for the SMTP connection
}

//SMTP command handler
//Takes in input from the session handler and sends the response string

//Interface into honeypot handler

//Error handler
func errorHandler(err error) {
	if err != nil {
		fmt.Print("Something went wrong with the SMTP connection handler")
		os.Exit(1)
	}
}
