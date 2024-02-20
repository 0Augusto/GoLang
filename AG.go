package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Indivíduo representa uma solução candidata
type Individuo struct {
	Genes []int // Genes representam as características da solução
}

// População representa uma coleção de indivíduos
type Populacao struct {
	Individuos []Individuo // Individuos é uma fatia de indivíduos
}

// Função de adaptação: calcula o lucro total da solução
func (i Individuo) Adaptação() int {
	total := 0
	for _, gene := range i.Genes {
		total += gene // Supõe-se que o lucro seja diretamente proporcional ao valor do gene
	}
	return total
}

// Função de crossover: combina genes de dois pais para gerar um filho
func Crossover(pai, mae Individuo) Individuo {
	filho := Individuo{}
	meio := len(pai.Genes) / 2
	filho.Genes = append(pai.Genes[:meio], mae.Genes[meio:]...)
	return filho
}

// Função de mutação: altera aleatoriamente um gene
func Mutação(individuo Individuo) {
	indice := rand.Intn(len(individuo.Genes))
	individuo.Genes[indice] = rand.Intn(10) // Suponha que o gene varie de 0 a 9
}

// Função principal do algoritmo genético
func AlgoritmoGenetico(tamanhoPopulacao, numGeracoes int) {
	// Inicializa a população inicial com indivíduos aleatórios
	populacao := Populacao{}
	for i := 0; i < tamanhoPopulacao; i++ {
		individuo := Individuo{}
		for j := 0; j < 50; j++ { // Suponha que cada solução tenha 50 genes
			individuo.Genes = append(individuo.Genes, rand.Intn(10)) // Suponha que o gene varie de 0 a 9
		}
		populacao.Individuos = append(populacao.Individuos, individuo)
	}

	// Itera sobre o número especificado de gerações
	for geracao := 0; geracao < numGeracoes; geracao++ {
		// Calcula a adaptação de cada indivíduo na população
		adaptacoes := make(map[int]int)
		for indice, individuo := range populacao.Individuos {
			adaptacoes[indice] = individuo.Adaptação()
		}

		// Ordena os indivíduos com base em sua adaptação (do maior para o menor)
		sort.Slice(populacao.Individuos, func(i, j int) bool {
			return adaptacoes[i] > adaptacoes[j]
		})

		// Seleciona os melhores indivíduos para a próxima geração
		melhores := populacao.Individuos[:tamanhoPopulacao/2]

		// Crossover: gera novos indivíduos combinando os genes dos pais
		novaPopulacao := Populacao{}
		for i := 0; i < tamanhoPopulacao/2; i++ {
			pai := melhores[rand.Intn(len(melhores))]
			mae := melhores[rand.Intn(len(melhores))]
			filho := Crossover(pai, mae)
			novaPopulacao.Individuos = append(novaPopulacao.Individuos, filho)
		}

		// Mutação: altera aleatoriamente alguns genes na nova população
		for _, individuo := range novaPopulacao.Individuos {
			if rand.Float32() < 0.01 { // Probabilidade de mutação de 25% é alto
				Mutação(individuo)
			}
		}

		// Substitui a população antiga pela nova população
		populacao.Individuos = novaPopulacao.Individuos
	}

	// Calcula a adaptação de cada indivíduo na população final
	adaptacoes := make(map[int]int)
	for indice, individuo := range populacao.Individuos {
		adaptacoes[indice] = individuo.Adaptação()
	}

	// Encontra e imprime a solução com o maior lucro
	melhorIndividuo := populacao.Individuos[0]
	melhorAdaptacao := adaptacoes[0]
	for i := 1; i < len(populacao.Individuos); i++ {
		if adaptacoes[i] > melhorAdaptacao {
			melhorIndividuo = populacao.Individuos[i]
			melhorAdaptacao = adaptacoes[i]
		}
	}
	fmt.Printf("Melhor solução encontrada: %+v\n", melhorIndividuo)
	fmt.Printf("Lucro máximo: %d\n", melhorAdaptacao)
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed aleatória

	// Parâmetros do algoritmo genético
	tamanhoPopulacao := 1000
	numGeracoes := 1000

	// Executa o algoritmo genético
	AlgoritmoGenetico(tamanhoPopulacao, numGeracoes)
}


