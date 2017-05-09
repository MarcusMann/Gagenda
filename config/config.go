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

// Procura por contatos
func procura(nome string, lista []byte) (contatos []string) {
	// expressão regular que verifica se o contato existe e pega eles.
	re := regexp.MustCompile("(?P<name>" + nome + ").+")

	// Se for != "" retorna pelo contato procurado mas se for == "" retorna todos os contatos.
	contatos = re.FindAllString(string(lista), -1)

	return
}

// EncontrarContatos procura por contatos
func EncontrarContatos() {
	// pergunta o nome do usuário
	questoes := Perguntas("Digite um nome para encontrar um contato cadastrado, ou pressione enter para exibir todos os contatos")

	// abre o arquivo de contatos
	listaContatos, err := ioutil.ReadFile(filename)

	// verifica se há algum erro ao tentar usar a expressão regular
	errMessage(err)

	// cria um slice de contatos para armazenar cada contato
	var contatos []string

	// Cria uma tabela.
	table := termtables.CreateTable()

	// Define os th (headers) da tabela
	table.AddHeaders("Nome", "Telefone")

	for _, cadaContato := range procura(questoes[0], listaContatos) {
		contatos = strings.Split(cadaContato, "  ") // faz um split nas strings, separando cada linha por dois espaços e atribuindo a um slice.
		// adiciona as linhas da tabela, tr
		table.AddRow(strings.Title(contatos[0]), contatos[1])
	}

	// mostra a tabela redenrizada
	fmt.Println(table.Render())
}

// CriarContato cria o arquivo para cada contato
func CriarContato() {
	questoes := Perguntas("Digite um nome: ", "Digite um número: ")

	contatos, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)

	// trata o erro
	errMessage(err)

	// fecha o documento depois que a função terminar
	defer contatos.Close()

	// escreve dentro do arquivo
	// Converte o número inteiro(idade) em string
	if _, err := contatos.WriteString(strings.ToLower(questoes[0]) + "  " + questoes[1] + "\n"); err != nil {
		panic(err)
	}

	fmt.Printf("Contato : %s salvo com sucesso! \n", questoes[0])
}

func errMessage(err error) {
	if err != nil {
		panic(err)
	}
}
