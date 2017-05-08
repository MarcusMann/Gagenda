package main

import (
	"os"

	"exercicios/gagenda/files"

	"github.com/paulrademacher/climenu"
)

func main() {
	menu := climenu.NewButtonMenu("Bem-Vindo a sua lista de contatos", "O que você gostaria de fazer?")
	menu.AddMenuItem("Cadastrar um Contato", "cadastrar")
	menu.AddMenuItem("Mostrar um Contato", "unico")
	menu.AddMenuItem("Mostrar todos os contatos", "todos")
	menu.AddMenuItem("Renomeia contato", "renomeia")

	action, scaped := menu.Run()
	if scaped {
		os.Exit(1)
	}

	switch action {
	case "cadastrar":
		files.CreateContact()
	case "unico":
		files.ShowUniqueContact()
	case "todos":
		files.ShowAllContacts()
	case "renomeia":
		files.RenameContacts()
	}
}
