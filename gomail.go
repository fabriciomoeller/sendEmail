package main

import (
	"log"

	"gopkg.in/gomail.v2"
)

func main() {
	// Configuração do servidor SMTP e porta
	smtpHost := "sandbox.smtp.mailtrap.io"
	smtpPort := 2525
	username := "1b4847980c6087"
	password := "2f4567c2b07c26"

	from := "teste@datainfo.inf.br"

	// Configuração do destinatário e mensagem
	to := "fabricio.moeller@datainfo.inf.br"
	subject := "Assunto do e-mail"
	body := "Email enviado hoje"

	// Configuração do objeto mensagem
	message := gomail.NewMessage()
	message.SetHeader("From", from)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/plain", body)

	// Configuração da autenticação
	d := gomail.NewDialer(smtpHost, smtpPort, username, password)
	//d.TLSConfig = &tls.Config{ServerName: "DATAEX19.datainfo.inf.br"}

	// Tenta enviar o e-mail
	if err := d.DialAndSend(message); err != nil {
		log.Fatalf("Falha ao enviar o e-mail: %v", err)
	}

	log.Println("E-mail enviado com sucesso!")
}
