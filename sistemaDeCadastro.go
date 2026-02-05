package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Leitor global
var reader = bufio.NewReader(os.Stdin)

// 1 - Imprimir uma mensagem de boas-vindas
func welcome() {
	fmt.Println("Bem-vindo ao sistema de filmes!")
}

// Função auxiliar para ler string
func readString() string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Função auxiliar para ler int
func readInt() int {
	value, _ := strconv.Atoi(readString())
	return value
}

// Função auxiliar para ler float64
func readFloat() float64 {
	value, _ := strconv.ParseFloat(readString(), 64)
	return value
}

// 2 - Função para cadastrar um filme
func createMovie() {
	fmt.Println("Digite o nome do filme:")
	name := readString()

	fmt.Println("Digite o ano de lançamento do filme:")
	yearRelease := readInt()

	fmt.Println("Digite o preço do filme:")
	moviePrice := readFloat()

	fmt.Printf(
		"O filme cadastrado é %s lançado em %d, ele custa - R$ %.2f\n",
		name,
		yearRelease,
		moviePrice,
	)
}

// 3 - Função para calcular a média de notas
func calculateAverage() float64 {
	fmt.Println("Digite quantas avaliações deseja fazer para o filme:")
	numRatings := readInt()

	var total float64

	for i := 0; i < numRatings; i++ {
		fmt.Printf("Digite a nota %d: ", i+1)
		note := readFloat()
		total += note
	}

	if numRatings == 0 {
		return 0
	}

	return total / float64(numRatings)
}

func main() {
	welcome()

	fmt.Println("Agora iremos cadastrar um filme no sistema.")
	createMovie()

	fmt.Println("Agora atribua as notas para cálculo da média de avaliações:")
	media := calculateAverage()

	fmt.Printf("A média das avaliações é: %.2f\n", media)
}
