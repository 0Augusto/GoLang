package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Função para calcular a sequência de Collatz para um número dado
func collatz(n uint64) ([]uint64, uint64) {
	steps := uint64(0)
	sequence := []uint64{}
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		steps++
		sequence = append(sequence, n)
	}
	return sequence, steps
}

// Função para imprimir um número em suas representações decimal, hexadecimal e octal
func printNumbers(num float64) {
	fmt.Printf("Decimal: %f\n", num)
	fmt.Printf("Hexadecimal: %X\n", uint64(num))
	fmt.Printf("Octal: %o\n", uint64(num))
}

// Função para imprimir um número em sua representação binária
func printBinary(num float64) {
	bits := math.Float64bits(num)
	fmt.Printf("%064b\n", bits)
}

func main() {
	// Criar um scanner para ler a entrada do usuário
	scanner := bufio.NewScanner(os.Stdin)

	// Solicita ao usuário inserir o primeiro número
	fmt.Println("Insira o primeiro número em formato binário:")
	if !scanner.Scan() {
		fmt.Println("Erro ao ler a entrada.")
		return
	}
	num1Input := scanner.Text()

	// Solicita ao usuário inserir o segundo número
	fmt.Println("Insira o segundo número em formato binário:")
	if !scanner.Scan() {
		fmt.Println("Erro ao ler a entrada.")
		return
	}
	num2Input := scanner.Text()

	// Converte as entradas do usuário para uint64
	num1Bin, err := strconv.ParseUint(strings.TrimSpace(num1Input), 2, 64)
	if err != nil {
		fmt.Println("Erro ao converter o primeiro número:", err)
		return
	}

	num2Bin, err := strconv.ParseUint(strings.TrimSpace(num2Input), 2, 64)
	if err != nil {
		fmt.Println("Erro ao converter o segundo número:", err)
		return
	}

	// Converte os números binários para float64
	num1 := float64(num1Bin)
	num2 := float64(num2Bin)

	// Operações
	sum := num1 + num2
	diff := num1 - num2
	product := num1 * num2
	quotient := num1 / num2

	// Imprimindo os resultados
	fmt.Println("Soma:")
	printNumbers(sum)
	fmt.Printf("Representação binária: ")
	printBinary(sum)
	fmt.Println("Subtração:")
	printNumbers(diff)
	fmt.Printf("Representação binária: ")
	printBinary(diff)
	fmt.Println("Multiplicação:")
	printNumbers(product)
	fmt.Printf("Representação binária: ")
	printBinary(product)
	fmt.Println("Divisão:")
	printNumbers(quotient)
	fmt.Printf("Representação binária: ")
	printBinary(quotient)

	// Calculando e exibindo a sequência de Collatz para os números inseridos pelo usuário
	fmt.Println("Sequência de Collatz para o primeiro número:")
	seq1, steps1 := collatz(uint64(num1))
	fmt.Println("Número:", uint64(num1))
	fmt.Println("Passos:", steps1)
	fmt.Println("Sequência:", seq1)

	fmt.Println("Sequência de Collatz para o segundo número:")
	seq2, steps2 := collatz(uint64(num2))
	fmt.Println("Número:", uint64(num2))
	fmt.Println("Passos:", steps2)
	fmt.Println("Sequência:", seq2)
}

