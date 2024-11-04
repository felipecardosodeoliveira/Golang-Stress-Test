package main

import (
	"flag"
	"fmt"
	"github/felipecardosodeoliveira/golang-stress-test/internal/tester"
	"os"
	"time"
)

func main() {
	url := flag.String("url", "", "URL do serviço a ser testado")
	requests := flag.Int("requests", 100, "Número total de requests")
	concurrency := flag.Int("concurrency", 10, "Número de chamadas simultâneas")
	flag.Parse()
	if *url == "" {
		fmt.Println("A URL deve ser fornecida")
		os.Exit(1)
	}
	startTime := time.Now()
	results := tester.RunLoadTest(*url, *requests, *concurrency)
	duration := time.Since(startTime)
	tester.GenerateReport(results, duration)
}
