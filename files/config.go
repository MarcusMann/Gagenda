package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/apcera/termtables"

	"bufio"
)

// EncontrarContatos procura por contatos
func EncontrarContatos(nome string) {
	// abre o arquivo de contatos
	listaContatos, err := ioutil.ReadFile("contacts/contatos.txt")
	errMessage(err) // verifica se há algum erro ao tentar usar a expressão regular

	// expressão regular que verifica se o contato existe e pega eles.
	re := regexp.MustCompile("(?P<name>" + nome + ").+")

	// Se for != "" retorna pelo contato procurado mas se for == "" retorna todos os contatos.
	procura := re.FindAllString(string(listaContatos), -1)

	var contatos []string // cria um slice de contatos para armazenar cada contato

	// Cria uma tabela.
	table := termtables.CreateTable()
	// Define os th (headers) da tabela
	table.AddHeaders("Nome", "Telefone")

	for _, cadaContato := range procura {
		contatos = strings.Split(cadaContato, "  ") // faz um split nas strings, separando cada linha por dois espaços e atribuindo a um slice.
		// adiciona as linhas da tabela, tr
		table.AddRow(contatos[0], contatos[1])
	}

	fmt.Println(table.Render())
}

// CreateContact cria o arquivo para cada contato
func CreateContact() {
	questoes := []string{"Digite um nome: ", "Digite um número: "}
	input := bufio.NewScanner(os.Stdin)
	var data []string

	for i := 0; i < len(questoes); i++ {
		fmt.Println(questoes[i])
		input.Scan()

		data = append(data, input.Text())
	}

	f, err := os.Create("contacts/pessoas.txt")

	errMessage(err)

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
