package main

import (
	"fmt"
	"log"
	"os"
	"sales-project/cmd/config"
	"sales-project/internal/ports/repositories/sales"
	"sales-project/pkg/csvreader"
)

func main() {
	cfg, err := config.FromEnv()
	if err != nil {
		panic(fmt.Errorf("can't read config: %w", err))
	}

	log.Printf("Running with config: %v", cfg)

	file, _ := os.Open(fmt.Sprintf("./files/%s.csv", cfg.FileName))
	defer file.Close()

	rows := csvreader.Read(file)
	sales.SaveSales(rows, cfg)
}
