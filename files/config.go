package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"bufio"

	"github.com/apcera/termtables"
)

// Contains verifica se uma string ta dentro de um []string
func Contains(contact string, s []string) bool {
	for _, contacts := range s {
		if contacts == contact {
			return true
		}
	}
	return false
}

// CreateContact cria o arquivo para cada contato
func CreateContact() {
	questoes := []string{"Digite seu nome: ", "Digite o seu número: "}
	input := bufio.NewScanner(os.Stdin)
	var data []string

	for i := 0; i < len(questoes); i++ {
		fmt.Println(questoes[i])
		input.Scan()

		data = append(data, input.Text())
	}

	f, err := os.Create("contacts/" + data[0] + ".txt")

	if err != nil {
		panic(err)
	}

	// fecha o documento depois que a função terminar
	defer f.Close()

	// escreve dentro do arquivo
	// Converte o número inteiro(idade) em string
	f.WriteString(data[0] + "\t" + data[1])

	// Caso não ocorra um erro, o arquivo irá salvar normalmente.
	fmt.Println(f.Name() + " criado com sucesso!")
}

// ShowUniqueContact Ler cada contato separadamente
func ShowUniqueContact() {
	input := bufio.NewScanner(os.Stdin)
	var nome string

	for i := 0; i < 1; i++ {
		fmt.Println("Digite o nome do contato: ")
		input.Scan()
		nome = input.Text()
	}

	f, err := ioutil.ReadFile("contacts/" + nome + ".txt")

	// verifica se ao ler o arquivo resulta em algum erro
	if err != nil {
		panic(err)
	}

	// remove a tabulação + espaço em branco dentro do arquivo e armazena tudo que tem lá dentro da variável contato
	contact := strings.Split(string(f), "\t")

	// Tabelas ...
	table := termtables.CreateTable()
	table.AddHeaders("Nome", "Telefone")
	table.AddRow(contact[0], contact[1])

	fmt.Println(table.Render())
}

// ShowAllContacts Mostra todos os contatos
func ShowAllContacts() {
	var contact []string

	// Tabelas ...
	table := termtables.CreateTable()
	table.AddHeaders("Nome", "Telefone")

	for _, filenames := range GetAllContacts() {
		f, err := ioutil.ReadFile("contacts/" + filenames)

		if err != nil {
			panic(err)
		}

		contact = strings.Split(string(f), "\t")
		table.AddRow(contact[0], contact[1])

	}

	fmt.Println(table.Render())

}

// GetAllContacts Mostra todos os contatos
func GetAllContacts() (contacts []string) {
	f, err := ioutil.ReadDir("contacts/")

	if err != nil {
		panic(err)
	}

	for _, contact := range f {
		contacts = append(contacts, contact.Name())
	}

	return
}

// RenameContacts renomeia o contato
func RenameContacts() {
	questoes := [2]string{"nome atual: ", "novo nome: "}
	var decisao []string

	input := bufio.NewScanner(os.Stdin)

	for i := 0; i < len(questoes); i++ {
		fmt.Println(questoes[i])
		input.Scan()
		decisao = append(decisao, input.Text())
	}

	rename := os.Rename("contacts/"+decisao[0]+".txt", "contacts/"+decisao[1]+".txt")

	if rename != nil {
		fmt.Printf("Ocorreu um erro ao tentar renomear o contato, tente novamente! erro: %v", rename)
	}
}

// DeleteContacts deleta um contato
func DeleteContacts() {
	fmt.Println("Digite o nome do contato que você gostaria de deletar: ")
	input := bufio.NewScanner(os.Stdin)

	delete := os.Remove("contacts/" + input.Text() + ".txt")

	if delete != nil {
		fmt.Printf("Ocorreu um erro ao tentar renomear o contato, tente novamente! erro: %v", delete)
	}

}
