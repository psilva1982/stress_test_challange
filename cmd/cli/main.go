package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

func main() {
	var url string
	var totalRequests int
	var concurrency int

	var rootCmd = &cobra.Command{
		Use:   "loadtest",
		Short: "CLI para testes de carga em serviços web",
		Run: func(cmd *cobra.Command, args []string) {
			loadTest(url, totalRequests, concurrency)
		},
	}

	rootCmd.Flags().StringVarP(&url, "url", "u", "", "URL do serviço a ser testado")
	rootCmd.Flags().IntVarP(&totalRequests, "requests", "r", 0, "Número total de requests")
	rootCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 1, "Número de chamadas simultâneas")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func loadTest(url string, totalRequests int, concurrency int) {
	if url == "" || totalRequests <= 0 || concurrency <= 0 {
		fmt.Println("Por favor, forneça a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.")
		return
	}

	fmt.Printf("Iniciando teste de carga para %s com %d requests e %d chamadas simultâneas...\n", url, totalRequests, concurrency)

	startTime := time.Now()
	var wg sync.WaitGroup

	successfulRequests := 0
	statusCodes := make(map[int]int)

	for concurrencyCounter := 0; concurrencyCounter < concurrency; concurrencyCounter++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for reqCounter := 0; reqCounter < totalRequests; reqCounter++ {
				statusCode := makeRequest(url)
				if statusCode == http.StatusOK {
					successfulRequests++
				}
				statusCodes[statusCode]++
			}
		}()
		time.Sleep(time.Duration(1000/concurrency) * time.Millisecond)
	}

	wg.Wait()

	elapsedTime := time.Since(startTime)

	fmt.Println("Relatório de Teste de Carga:")
	fmt.Printf("Tempo total gasto na execução: %s\n", elapsedTime)
	fmt.Printf("Quantidade total de requests realizados: %d\n", totalRequests*concurrency)
	fmt.Printf("Quantidade de requests com status HTTP 200: %d\n", successfulRequests)
	fmt.Println("Distribuição de outros códigos de status HTTP:")
	for code, count := range statusCodes {
		fmt.Printf("  - Status %d: %d\n", code, count)
	}
}

func makeRequest(url string) int {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erro ao fazer a requisição:", err)
		return http.StatusInternalServerError
	}

	defer resp.Body.Close()
	return resp.StatusCode
}
