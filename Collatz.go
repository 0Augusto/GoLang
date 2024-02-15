package main

import (
	"fmt"
)

// Função para calcular a conjectura de Collatz para um determinado número
func collatz(n int) []int {
	sequence := []int{n}

	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		sequence = append(sequence, n)
	}

	return sequence
}

func main() {
	var num int

	// Solicitar ao usuário que insira um número
	fmt.Print("Insira um número inteiro positivo: ")
	_, err := fmt.Scan(&num)
	if err != nil || num <= 0 {
		fmt.Println("Por favor, insira um número inteiro positivo válido.")
		return
	}

	// Calcular a conjectura de Collatz para o número inserido
	sequence := collatz(num)

	// Imprimir a sequência
	fmt.Printf("Conjectura de Collatz para %d: ", num)
	for i, val := range sequence {
		if i > 0 {
			fmt.Print(" -> ")
		}
		fmt.Print(val)
	}
	fmt.Println()
}

