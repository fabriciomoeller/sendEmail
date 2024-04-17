package main

import (
	"log"

	"gopkg.in/gomail.v2"
)

func sendEmail(from string, to string, subject string, body string) (string, error) {
	// Configuração do servidor SMTP e porta
	// Configuração do servidor SMTP e porta
	smtpHost := "sandbox.smtp.mailtrap.io"
	smtpPort := 2525
	username := "1b4847980c6087"
	password := "2f4567c2b07c26"

	// Configuração do objeto mensagem
	message := gomail.NewMessage()
	message.SetHeader("From", from)
	// Configuração do destinatário e mensagem
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/plain", body)

	// Configuração da autenticação
	d := gomail.NewDialer(smtpHost, smtpPort, username, password)
	//d.TLSConfig = &tls.Config{ServerName: "DATAEX19.datainfo.inf.br"}

	// Tenta enviar o e-mail
	if err := d.DialAndSend(message); err != nil {
		return "Falha ao enviar o e-mail: %v", err
	} else {
		return "E-mail enviado com sucesso!", nil
	}

}

func main() {

	sucesso, err := sendEmail("teste@datainfo.inf.br", "fabricio.moeller@datainfo.inf.br", "Assunto do e-mail", "Email enviado hoje")

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(sucesso)
	}

}
