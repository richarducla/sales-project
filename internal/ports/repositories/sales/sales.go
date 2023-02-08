package sales

import (
	"context"
	"fmt"
	"sales-project/cmd/config"
	"sales-project/internal/models"
	"sales-project/pkg/postgres"
	"time"

	"github.com/jackc/pgx/v4"
)

const (
	deleteQuery = "delete from sales"
)

func RemoveAll(cfg config.Config) error {
	conn, err := postgres.NewConn(cfg)
	if err != nil {
		fmt.Println("newDB", err)
		return err
	}

	defer func() {
		_ = conn.Close(context.Background())
		fmt.Println("db closed")
	}()

	_, err = conn.Query(context.Background(), deleteQuery)
	if err != nil {
		return err
	}

	return nil

}

func SaveSales(rows []models.Sale, cfg config.Config) error {
	chunks := getChunks(rows, cfg.ChunkSize)

	countChunks := len(chunks)
	finishChan := make(chan int)

	ti := time.Now()
	for i := 0; i < countChunks; i++ {
		go func(c chan int, i int, cfg config.Config) {
			if err := insertSales(chunks[i], cfg); err != nil {
				fmt.Println("Error", err)
			}
			c <- 1
		}(finishChan, i, cfg)
	}

	finishedGophers := 0
	finishLoop := false
	for {
		if finishLoop {
			break
		}
		select {
		case n := <-finishChan:
			finishedGophers += n
			if finishedGophers == countChunks {
				finishLoop = true
			}
		}
	}

	fmt.Println("Time in insertion: ", time.Since(ti))
	return nil
}

func insertSales(rows []models.Sale, cfg config.Config) error {
	conn, err := postgres.NewConn(cfg)
	if err != nil {
		fmt.Println("newDB", err)
		return err
	}

	defer func() {
		_ = conn.Close(context.Background())
		fmt.Println("db closed")
	}()

	count, err := conn.CopyFrom(
		context.Background(),
		pgx.Identifier{"sales"},
		[]string{"point_of_sale", "product", "date", "stock"},
		pgx.CopyFromSlice(len(rows), func(i int) ([]interface{}, error) {
			return []interface{}{rows[i].PointOfSale, rows[i].Product, rows[i].Date, rows[i].Stock}, nil
		}),
	)
	if err != nil {
		return fmt.Errorf("conn.CopyFrom %w", err)
	}

	fmt.Println("Number of records inserted", count)

	return nil
}

func getChunks(rows []models.Sale, chunkSize int) [][]models.Sale {
	var divided [][]models.Sale
	for i := 0; i < len(rows); i += chunkSize {
		end := i + chunkSize

		if end > len(rows) {
			end = len(rows)
		}

		divided = append(divided, rows[i:end])
	}
	return divided
}
