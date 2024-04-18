package main

import (
	"testing"
)

// Variáveis de teste. Pode ser alterado para testar diferentes cenários.
var tests = []struct {
	to      string
	from    string
	subject string
	body    string
	error   string
}{
	{
		to:      "fabricio@envioemail.com",
		from:    "fabricio.moeller@datainfo.inf.br",
		subject: "Teste de email com sucesso",
		body:    "Email enviado por teste unitário",
	},
	{
		to:      "",
		from:    "fabricio.moeller@datainfo.inf.br",
		subject: "Teste de email com erro no TO",
		body:    "Email enviado por teste unitário",
		error:   "Erro esperado no TO",
	},
	{
		to:      "fabricio@envioemail.com",
		from:    "",
		subject: "Teste de email com erro no FROM",
		body:    "Email enviado por teste unitário",
		error:   "Erro esperado no FROM",
	},
}

// TestSendEmail testa a função sendEmail.

func TestSendEmail(t *testing.T) {

	for _, test := range tests {
		if (test.to == "" || test.from == "") && testing.Short() {
			t.Skip("Ignorando teste")
		} else {
			t.Run(test.subject, func(t *testing.T) {
				t.Parallel()
				sucesso, err := sendEmail(
					test.from,
					test.to,
					test.subject,
					test.body)
				if err != nil {
					if err.Error() != `gomail: could not send email 1: gomail: invalid address "": mail: no address` {
						t.Fatal("Gerou o ERRO")
					} else {
						t.Log(test.error, err)
					}
				} else {
					t.Log(sucesso)
				}
			})
		}
	}
}
