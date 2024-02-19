package main

import (
	"fmt"
	"net"
	"sort"
)

// worker é uma função que é executada concorrentemente para verificar a abertura de portas.
// Ele recebe canais de entrada (ports) e de saída (results) para coordenar as operações.
func worker(ports, results chan int) {
	// Loop sobre os valores recebidos do canal 'ports'
	for p := range ports {
		// Constrói o endereço de IP e porta na forma de uma string
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		
		// Tenta estabelecer uma conexão TCP com o endereço
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// Se houver um erro na conexão, envia 0 para o canal de resultados e continua para a próxima porta
			results <- 0
			continue
		}
		// Se a conexão for bem-sucedida, fecha-a
		conn.Close()
		// Envia o número da porta para o canal de resultados
		results <- p
	}
}

func main() {
	// Criação de canais para coordenar a comunicação entre goroutines
	ports := make(chan int, 100) // Canal para enviar números de porta para serem escaneados
	results := make(chan int)     // Canal para receber os resultados do escaneamento
	var openports []int           // Slice para armazenar as portas abertas
	
	// Inicia goroutines trabalhadoras
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	// Goroutine anônima para enviar as portas para serem escaneadas
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i // Envia o número da porta para o canal 'ports'
		}
	}()

	// Recebe os resultados do canal de resultados e armazena as portas abertas
	for i := 0; i < 1024; i++ {
		port := <-results // Recebe um resultado do canal 'results'
		if port != 0 {
			openports = append(openports, port) // Se a porta for diferente de 0, ela é considerada aberta
		}
	}

	// Fecha os canais após a conclusão do escaneamento
	close(ports)
	close(results)

	// Ordena as portas abertas
	sort.Ints(openports)

	// Imprime as portas abertas
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}

