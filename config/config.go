package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/apcera/termtables"

	"bufio"
)

const (
	// arquivo de contatos
	filename = "contatos.txt"
)

// Perguntas cria perguntas e retorna as respostas das perguntas que foram criadas
// Perguntas são passadas como argumentos, separadas por ,
func Perguntas(perguntas ...string) (respostas []string) {
	input := bufio.NewScanner(os.Stdin)

	for i := 0; i < len(perguntas); i++ {
		fmt.Println(perguntas[i])

		input.Scan()

		respostas = append(respostas, input.Text())
	}

	return
}

// EncontrarContatos procura por contatos
func EncontrarContatos(nome string) {
	// abre o arquivo de contatos
	listaContatos, err := ioutil.ReadFile(filename)
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
	questoes := Perguntas("Digite um nome: ", "Digite um número: ")

	contatos, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)

	// trata o erro
	errMessage(err)

	// fecha o documento depois que a função terminar
	defer contatos.Close()

	// escreve dentro do arquivo
	// Converte o número inteiro(idade) em string
	if _, err := contatos.WriteString(questoes[0] + "  " + questoes[1] + "\n"); err != nil {
		panic(err)
	}

	fmt.Printf("Contato : %s salvo com sucesso! \n", questoes[0])
}

// RenomeiaContato renomeia o contato
func RenomeiaContato() {
	questoes := Perguntas("antigo nome: ", "novo nome: ")

	renomeia := os.Rename("contacts/"+questoes[0]+".txt", "contacts/"+questoes[1]+".txt")

	if renomeia != nil {
		fmt.Printf("Ocorreu um erro ao tentar renomear o contato, tente novamente! erro: %v", renomeia)
	}
}

// DeletaContato deleta um contato
func DeletaContato() {
	questoes := Perguntas("Digite o nome do contato que você gostaria de deletar: ")

	delete := os.Remove("contacts/" + questoes[0] + ".txt")

	// trata o erro
	errMessage(delete)
}

func errMessage(err error) {
	if err != nil {
		panic(err)
	}
}
