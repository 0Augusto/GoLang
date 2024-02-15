package main

import (
	"container/list"
	"fmt"
)

// Função para calcular a conjectura de Collatz para um determinado número
func collatz(n int) int {
	if n%2 == 0 {
		return n / 2
	} else {
		return 3*n + 1
	}
}

// Estrutura de dados para armazenar informações na cache
type Cache struct {
	data      map[int]int // Mapa para armazenar números e seus resultados de Collatz
	order     *list.List   // Lista duplamente encadeada para manter a ordem temporal dos números na cache
	capacity  int          // Capacidade máxima da cache
}

// Função para criar uma nova cache com capacidade especificada
func NewCache(capacity int) *Cache {
	return &Cache{
		data:     make(map[int]int),
		order:    list.New(),
		capacity: capacity,
	}
}

// Função para obter o resultado de Collatz de um número da cache
func (c *Cache) Get(n int) int {
	// Verifica se o número já está na cache
	if val, ok := c.data[n]; ok {
		// Move o número para o final da lista (indicando que foi acessado recentemente)
		for e := c.order.Front(); e != nil; e = e.Next() {
			if e.Value.(int) == n {
				c.order.MoveToBack(e)
				break
			}
		}
		return val
	}

	// Se o número não estiver na cache, calcula seu resultado de Collatz
	result := collatz(n)

	// Verifica se a cache atingiu sua capacidade máxima
	if len(c.data) >= c.capacity {
		// Remove o número mais antigo da cache (o primeiro da lista)
		delete(c.data, c.order.Front().Value.(int))
		c.order.Remove(c.order.Front())
	}

	// Adiciona o novo número à cache e à lista (indicando que foi acessado recentemente)
	c.data[n] = result
	c.order.PushBack(n)

	return result
}

func main() {
	// Criar uma nova cache com capacidade de 5
	cache := NewCache(5)

	// Exemplo de uso da cache com a conjectura de Collatz
	numbers := []int{10, 20, 5, 16, 8, 10, 13, 20}

	for _, num := range numbers {
		fmt.Printf("Resultado de Collatz para %d: %d\n", num, cache.Get(num))
	}
}

