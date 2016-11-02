package main

import (
	"encoding/csv"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"sync"
)

var (
	minlat float64 = -90
	maxlat float64 = 90
	minlon float64 = -180
	maxlon float64 = 180
)

func main() {
	// Set up the pipe to write data directly into the Reader.
	pr, pw := io.Pipe()
	// Write JSON-encoded data to the Writer end of the pipe.
	// Write in a separate concurrent goroutine, and remember
	// to Close the PipeWriter, to signal to the paired PipeReader
	// that we’re done writing.
	go func() {

		var wg sync.WaitGroup
		ch := make(chan []string)

		for i := 0; i <= runtime.NumCPU(); i++ {
			wg.Add(1)
			go worker(ch, &wg)
		}

		go func() {
			wg.Wait()
			close(ch)
		}()

		w := csv.NewWriter(pw)
		for record := range ch {
			err := w.Write(record)

			if err != nil {
				log.Fatal(err)
			}
		}

		w.Flush()
		pw.Close()

	}()
	// Send the HTTP request. Whatever is read from the Reader
	// will be sent in the request body.
	// As data is written to the Writer, it will be available
	// to read from the Reader.
	log.Println(url())
	_, err := http.Post(url(), "text/csv", pr)

	if err != nil {
		log.Fatal(err)
	}

}

func streamAmountForWorker() int {
	cpuCoresCount := runtime.NumCPU()
	return streamingAmount() / cpuCoresCount
}

func streamingAmount() int {

	if len(os.Args) >= 3 && os.Args[2] != "" {
		i64, err := strconv.Atoi(os.Args[2])
		if err != nil {
			panic("OMG! invalid csv stream amount given!")
		}
		return i64
	}

	return 1000000000

}

func url() string {
	if len(os.Args) >= 2 && os.Args[1] != "" {
		return os.Args[1]
	} else {
		return "http://localhost:8080"
	}
}

func f64ToString(f float64) string {
	return strconv.FormatFloat(f, 'E', -1, 64)
}

func getID() string {
	return "X" + strconv.Itoa(rand.Intn(100))
}

func random(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

func worker(ch chan<- []string, wg *sync.WaitGroup) {
	amount := streamAmountForWorker()
	for i := 0; i <= amount; i++ {
		v := make([]string, 3)
		v[0] = getID()
		v[1] = f64ToString(random(minlat, maxlat))
		v[2] = f64ToString(random(minlon, maxlon))
		ch <- v
	}
	wg.Done()
}
