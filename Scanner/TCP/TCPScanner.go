package main

import (
	"fmt"
	"net"
)

func main() {
	// Loop de 1 a 1024 para verificar as primeiras 1024 portas
	for i := 1; i <= 1024; i++ {
		// Criação de uma goroutine para cada iteração do loop para execução concorrente
		go func(j int) {
			// Construção do endereço de IP e porta na forma de string
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			
			// Tentativa de conexão TCP com o endereço obtido
			conn, err := net.Dial("tcp", address)
			
			// Verifica se ocorreu um erro na conexão
			if err != nil {
				// Se ocorreu um erro, retorna silenciosamente (sem imprimir nada)
				return
			}
			
			// Após a tentativa de conexão bem-sucedida, fecha a conexão
			conn.Close()
			
			// Imprime a mensagem indicando que a porta está aberta
			fmt.Printf("%d open\n", j)
		}(i) // Passa o valor atual de 'i' para a função anônima como parâmetro 'j'
	}
}

