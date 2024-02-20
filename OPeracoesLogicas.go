package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Função para converter um número binário em uint64
func binaryToUint64(bin string) (uint64, error) {
	// Converte o número binário para uint64
	val, err := strconv.ParseUint(bin, 2, 64)
	if err != nil {
		return 0, err
	}
	return val, nil
}

// Função para converter uint64 para representações binária, decimal, hexadecimal e octal
func printNumbers(num uint64) {
	fmt.Printf("Decimal: %d\n", num)
	fmt.Printf("Hexadecimal: %X\n", num)
	fmt.Printf("Octal: %o\n", num)
}

// Função para realizar a operação lógica NAND entre dois números binários
func nand(num1, num2 uint64) uint64 {
	// A operação NAND retorna 0 apenas quando ambos os bits de entrada são 1
	return ^(num1 & num2)
}

// Função para realizar a operação lógica AND entre dois números binários
func and(num1, num2 uint64) uint64 {
	// A operação AND retorna 1 apenas quando ambos os bits de entrada são 1
	return num1 & num2
}

// Função para realizar a operação lógica OR entre dois números binários
func or(num1, num2 uint64) uint64 {
	// A operação OR retorna 1 se pelo menos um dos bits de entrada for 1
	return num1 | num2
}

// Função para realizar a operação lógica NOR entre dois números binários
func nor(num1, num2 uint64) uint64 {
	// A operação NOR retorna 1 apenas quando ambos os bits de entrada são 0
	return ^(num1 | num2)
}

// Função para realizar a operação lógica XOR entre dois números binários
func xor(num1, num2 uint64) uint64 {
	// A operação XOR retorna 1 apenas quando os bits de entrada são diferentes
	return num1 ^ num2
}

func main() {
	// Criar um scanner para ler a entrada do usuário
	scanner := bufio.NewScanner(os.Stdin)

	// Solicita ao usuário inserir o primeiro número em formato binário
	fmt.Println("Insira o primeiro número em formato binário:")
	if !scanner.Scan() {
		fmt.Println("Erro ao ler a entrada.")
		return
	}
	num1Input := scanner.Text()

	// Solicita ao usuário inserir o segundo número em formato binário
	fmt.Println("Insira o segundo número em formato binário:")
	if !scanner.Scan() {
		fmt.Println("Erro ao ler a entrada.")
		return
	}
	num2Input := scanner.Text()

	// Converte as entradas do usuário para uint64
	num1, err := binaryToUint64(strings.TrimSpace(num1Input))
	if err != nil {
		fmt.Println("Erro ao converter o primeiro número:", err)
		return
	}

	num2, err := binaryToUint64(strings.TrimSpace(num2Input))
	if err != nil {
		fmt.Println("Erro ao converter o segundo número:", err)
		return
	}

	// Realiza as operações lógicas entre os números inseridos pelo usuário
	fmt.Println("Operações lógicas:")
	fmt.Println("NAND:")
	fmt.Println("Resultado em binário:", strconv.FormatUint(nand(num1, num2), 2))
	fmt.Println("AND:")
	fmt.Println("Resultado em binário:", strconv.FormatUint(and(num1, num2), 2))
	fmt.Println("OR:")
	fmt.Println("Resultado em binário:", strconv.FormatUint(or(num1, num2), 2))
	fmt.Println("NOR:")
	fmt.Println("Resultado em binário:", strconv.FormatUint(nor(num1, num2), 2))
	fmt.Println("XOR:")
	fmt.Println("Resultado em binário:", strconv.FormatUint(xor(num1, num2), 2))
}

