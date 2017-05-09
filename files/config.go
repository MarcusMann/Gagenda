package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"bufio"
)

// // Tabelas ...
// table := termtables.CreateTable()
// table.AddHeaders("Nome", "Telefone")
// table.AddRow(contact[0], contact[1])

// fmt.Println(table.Render())

// Contains verifica se uma string ta dentro de um []string
func Contains(contact string, s []string) bool {
	for _, contacts := range s {
		if contacts == contact {
			return true
		}
	}
	return false
}

// FindContacts acha todos os contatos ou apenas um contato
func FindContacts(nome string) {
	listaContatos, err := ioutil.ReadFile("contacts/contatos.txt")

	// verifica se há algum erro ao tentar usar a expressão regular
	errMessage(err)

	// expressão regular que verifica se o contato existe e pega eles.
	procuraContato, err := regexp.Compile("(?P<name>" + nome + ").+")

	// verifica se há algum erro ao tentar usar a expressão regular
	errMessage(err)

	f := string(listaContatos)

	fmt.Println(procuraContato.FindString(f))

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

func errMessage(err error) {
	if err != nil {
		panic(err)
	}
}
