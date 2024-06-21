package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	url         string
	numRequests int
	concurrency int
)

func init() {
	flag.StringVar(&url, "url", "", "URL do serviço a ser testado")
	flag.IntVar(&numRequests, "requests", 1, "Número total de requests")
	flag.IntVar(&concurrency, "concurrency", 1, "Número de chamadas simultâneas")
}

type result struct {
	duration   time.Duration
	statusCode int
}

func makeRequest(wg *sync.WaitGroup, results chan<- result) {
	defer wg.Done()

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erro ao fazer request:", err)
		results <- result{duration: 0, statusCode: 0}
		return
	}
	defer resp.Body.Close()

	duration := time.Since(start)
	results <- result{duration: duration, statusCode: resp.StatusCode}
}

func main() {
	flag.Parse()

	if url == "" {
		fmt.Println("A URL é obrigatória")
		return
	}

	var wg sync.WaitGroup
	results := make(chan result, numRequests)
	startTime := time.Now()

	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go makeRequest(&wg, results)
		if (i+1)%concurrency == 0 {
			wg.Wait()
		}
	}

	wg.Wait()
	close(results)

	totalDuration := time.Since(startTime)
	successfulRequests := 0
	statusCounts := make(map[int]int)
	var totalResponseTime time.Duration

	for result := range results {
		if result.statusCode == 200 {
			successfulRequests++
		}
		if result.statusCode != 0 {
			statusCounts[result.statusCode]++
		}
		totalResponseTime += result.duration
	}

	fmt.Printf("Relatório:\n")
	fmt.Printf("Tempo total gasto na execução: %v\n", totalDuration)
	fmt.Printf("Quantidade total de requests realizados: %d\n", numRequests)
	fmt.Printf("Quantidade de requests com status HTTP 200: %d\n", successfulRequests)
	fmt.Printf("Distribuição de outros códigos de status HTTP:\n")
	for statusCode, count := range statusCounts {
		if statusCode != 200 {
			fmt.Printf("  %d: %d\n", statusCode, count)
		}
	}
}
