package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Função para imprimir um número em sua representação binária
func printBinary(num float64) {
	bits := math.Float64bits(num)
	fmt.Printf("%064b\n", bits)
}

func main() {
	// Solicita ao usuário inserir os números para realizar as operações
	var num1Input, num2Input string
	fmt.Println("Insira o primeiro número:")
	fmt.Scanln(&num1Input)

	fmt.Println("Insira o segundo número:")
	fmt.Scanln(&num2Input)

	// Converte as entradas do usuário para float64
	num1, err := strconv.ParseFloat(strings.TrimSpace(num1Input), 64)
	if err != nil {
		fmt.Println("Erro ao converter o primeiro número:", err)
		return
	}

	num2, err := strconv.ParseFloat(strings.TrimSpace(num2Input), 64)
	if err != nil {
		fmt.Println("Erro ao converter o segundo número:", err)
		return
	}

	// Operações
	sum := num1 + num2
	diff := num1 - num2
	product := num1 * num2
	quotient := num1 / num2

	// Imprimindo os resultados
	fmt.Println("Soma:")
	fmt.Printf("   %.2f\n", sum)
	fmt.Printf("   Representação binária: ")
	printBinary(sum)

	fmt.Println("Subtração:")
	fmt.Printf("   %.2f\n", diff)
	fmt.Printf("   Representação binária: ")
	printBinary(diff)

	fmt.Println("Multiplicação:")
	fmt.Printf("   %.2f\n", product)
	fmt.Printf("   Representação binária: ")
	printBinary(product)

	fmt.Println("Divisão:")
	fmt.Printf("   %.2f\n", quotient)
	fmt.Printf("   Representação binária: ")
	printBinary(quotient)
}

