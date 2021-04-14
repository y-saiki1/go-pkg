package mail

import (
	"fmt"
	"net/smtp"
	"os"
)

func Send(to, subject string, msgList ...string) error {
	var msg string
	for _, v := range msgList {
		msg += v
		msg += "\r\n"
	}
	addr := os.Getenv("SERVICE_MAIL_ADDRESS")
	pass := os.Getenv("SERVICE_MAIL_PASSWORD")
	host := os.Getenv("SERVICE_MAIL_HOST")
	port := os.Getenv("SERVICE_MAIL_PORT")
	fmt.Println(
		addr,
		pass,
		host,
		port,
	)

	auth := smtp.PlainAuth("", addr, pass, host)
	msgs := []byte("" +
		"From: iCashSupport <" + addr + ">\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + msg + "",
	)

	hostWithPort := host + ":" + port
	err := smtp.SendMail(hostWithPort, auth, addr, []string{to}, msgs)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
