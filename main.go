package main

import (
	"exercicios/gagenda/files"

	_ "github.com/paulrademacher/climenu"
)

func main() {
	// menu := climenu.NewButtonMenu("Bem-Vindo a sua lista de contatos", "O que você gostaria de fazer?")
	// menu.AddMenuItem("Cadastrar um Contato", "cadastrar")
	// menu.AddMenuItem("Mostrar um Contato", "unico")
	// menu.AddMenuItem("Mostrar todos os contatos", "todos")
	// menu.AddMenuItem("Renomeiar contato", "renomeiar")
	// menu.AddMenuItem("Deletar contato", "deletar")

	// action, scaped := menu.Run()
	// if scaped {
	// 	os.Exit(1)
	// }

	// switch action {
	// case "cadastrar":
	// 	files.CreateContact()
	// case "unico":
	// 	files.FindContacts("marcus")
	// case "renomeiar":
	// 	files.RenameContacts()
	// case "deletar":
	// 	files.DeleteContacts()
	// }
	files.EncontrarContatos("")
}
