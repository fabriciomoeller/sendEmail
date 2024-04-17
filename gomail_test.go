package main

import (
	"testing"
)

// TestSendEmail testa a função sendEmail.

func TestSendEmail(t *testing.T) {
	sucesso, err := sendEmail(
		"fabricio@envioemail.comr",
		"fabricio.moeller@datainfo.inf.br",
		"Teste de email 2",
		"Email enviado por teste unitário")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(sucesso)
	}
}

func TestSendEmailEmptyTo(t *testing.T) {
	sucesso, err := sendEmail(
		"",
		"fabricio.moeller@datainfo.inf.br",
		"Teste de email 2",
		"Email enviado por teste unitário")
	if err != nil {
		if err.Error() != `gomail: could not send email 1: gomail: invalid address "": mail: no address` {
			t.Fatal("Gerou o ERRO")
		} else {
			t.Log("Erro esperado no TO: ", err)
		}
	} else {
		t.Log(sucesso)
	}
}

func TestSendEmailEmptyFrom(t *testing.T) {
	sucesso, err := sendEmail(
		"fabricio@envioemail.com",
		"",
		"Teste de email 2",
		"Email enviado por teste unitário")
	if err != nil {
		if err.Error() != `gomail: could not send email 1: gomail: invalid address "": mail: no address` {
			t.Fatal(err)
		} else {
			t.Log("Erro esperado no FROM: ", err)
		}
	} else {
		t.Log(sucesso)
	}
}
