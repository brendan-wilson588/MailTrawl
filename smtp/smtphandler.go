package smtp

import (
	"fmt"
	"net"
	"os"
)

const (
	smtpAddr string = ":25"
)

type smtpSession struct {
	senderAddress          string
	recipientAddress       string
	data                   string
	clientReportedHostname string
	sourceAddress          string
}

//Main connection loop
func startSMTPServer() {
	//Start message
	tcpAddr, err := net.ResolveTCPAddr("ipv4", smtpAddr)
	//setup connection listener
	listener, err := net.ListenTCP("tcp", tcpAddr)
	errorHandler(err)

	//Main listen loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Something went wrong with the SMTP listener")
		}
		smtpSessionHandler(conn)
	}
}

func smtpSessionHandler(conn net.Conn) {
	//Get the input from the connection

	//Send HELO

	//Parse the first word for the SMTP connection
}

//SMTP command checker

//Interface into honeypot handler

//Error handler
func errorHandler(err error) {
	if err != nil {
		fmt.Print("Something went wrong with the SMTP connection handler")
		os.Exit(1)
	}
}
