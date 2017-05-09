package main

import (
	"exercicios/gagenda/config"
	"os"

	"github.com/paulrademacher/climenu"
)

func main() {
	menu := climenu.NewButtonMenu("Bem-Vindo a sua lista de contatos", "O que você gostaria de fazer?")
	menu.AddMenuItem("Cadastrar um Contato", "cadastrar")
	menu.AddMenuItem("Mostrar Contatos", "contatos")

	action, scaped := menu.Run()
	if scaped {
		os.Exit(1)
	}

	switch action {
	case "cadastrar":
		config.CriarContato()
	case "contatos":
		config.EncontrarContatos()
	}

}
