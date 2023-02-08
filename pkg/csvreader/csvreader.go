package csvreader

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sales-project/internal/models"
	"strconv"
	"strings"
	"sync"
	"time"
)

// with Worker pools
func Read(f *os.File) []models.Sale {
	sales := []models.Sale{}
	fcsv := csv.NewReader(f)
	numWps := 100
	jobs := make(chan []string, numWps)
	res := make(chan models.Sale)

	var wg sync.WaitGroup
	worker := func(jobs <-chan []string, results chan<- models.Sale) {
		for {
			select {
			case job, ok := <-jobs: // you must check for readable state of the channel.
				if !ok {
					return
				}
				sale, err := parseSale(job)
				if err == nil {
					results <- sale
				}
			}
		}
	}

	// init workers
	for w := 0; w < numWps; w++ {
		wg.Add(1)
		go func() {
			// this line will exec when chan `res` processed output at line 107 (func worker: line 71)
			defer wg.Done()
			worker(jobs, res)
		}()
	}

	go func() {
		for {
			rStr, err := fcsv.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("ERROR: ", err.Error())
				break
			}
			jobs <- rStr
		}
		close(jobs) // close jobs to signal workers that no more job are incoming.
	}()

	go func() {
		wg.Wait()
		close(res) // when you close(res) it breaks the below loop.
	}()

	for r := range res {
		sales = append(sales, r)
	}

	return sales
}

func parseSale(data []string) (models.Sale, error) {
	sale := strings.Split(data[0], ";")

	date, err := time.Parse("2006-01-02", sale[2])
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error during conversion in date")
		return models.Sale{}, err
	}

	stock, err := strconv.Atoi(sale[3])
	if err != nil {
		fmt.Println("Error during conversion in stock")
		return models.Sale{}, err
	}

	return models.Sale{
		PointOfSale: sale[0],
		Product:     sale[1],
		Date:        date,
		Stock:       stock,
	}, nil
}
