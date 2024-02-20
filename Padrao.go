package main

import "fmt"

// Função para aplicar a conjectura de Collatz a um número
func collatz(n uint64) uint64 {
    count := uint64(0)
    for n != 1 {
        if n%2 == 0 {
            n = n / 2
        } else {
            n = 3*n + 1
        }
        count++
    }
    return count
}

func main() {
    // Armazenamento do texto em um registrador
    var textRegister string
    textRegister = "A conjectura de Collatz é um problema matemático intrigante que ainda não foi completamente resolvido. Ela afirma que, começando de qualquer número inteiro positivo, o resultado sempre será 1 após um número finito de passos, seguindo duas regras simples de iteração: se o número for par, divida-o por 2; se for ímpar, multiplique-o por 3 e some 1."

    // Imprimir o texto armazenado no registrador
    fmt.Println("Texto armazenado no registrador:")
    fmt.Println(textRegister)

    // Aplicar a conjectura de Collatz a alguns números e imprimir os resultados
    fmt.Println("\nResultados da conjectura de Collatz para alguns números:")
    fmt.Println("Collatz(10):", collatz(10))
    fmt.Println("Collatz(27):", collatz(27))
    fmt.Println("Collatz(50):", collatz(50))
}

