package workloadgen

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	// "strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/liam0215/anarres/workflow/frontend"
)

// The WorkloadGen interface, which the Blueprint compiler will treat as a
// Workflow service
type SimpleWorkload interface {
	ImplementsSimpleWorkload(context.Context) error
}

// workloadGen implementation
type workloadGen struct {
	SimpleWorkload

	frontend frontend.Frontend
}

var sizeKb = flag.Int("sizeKb", 64, "Size of value in KB")
var numWorkers = flag.Int("numWorkers", 2, "Number of workers to send requests in parallel")
var duration = flag.String("duration", "15s", "Experiment duration (e.g. 15s)")

func NewSimpleWorkload(ctx context.Context, frontend frontend.Frontend) (SimpleWorkload, error) {
	return &workloadGen{frontend: frontend}, nil
}

func (s *workloadGen) Run(ctx context.Context) error {
	var reqCount uint64

	f, err := os.Open("xml")
	if err != nil {
		return fmt.Errorf("failed to open xml file: %w", err)
	}
	defer f.Close()

	limited := io.LimitReader(f, int64(*sizeKb)*1024)

	buf, err := io.ReadAll(limited)
	if err != nil {
		return fmt.Errorf("failed to read from xml: %w", err)
	}

	payload := string(buf)
	fmt.Printf("Using payload of size %d bytes\n", len(payload))

	stop := make(chan bool)

	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-stop:
				return
			case <-ticker.C:
				// swap out and reset the counter
				c := atomic.SwapUint64(&reqCount, 0)
				fmt.Printf("→ Throughput: %d req/s\n", c)
			}
		}
	}()

	dur, err := time.ParseDuration(*duration)
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	for id := 0; id < *numWorkers; id++ {
		key := fmt.Sprintf("key-%d", id)
		if err := s.frontend.Put(ctx, key, payload); err != nil {
			return fmt.Errorf("priming client definitions: %w", err)
		}

		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case <-stop:
					fmt.Printf("[worker %d] Stopping\n", workerID)
					return
				default:
					err := s.frontend.Put(ctx, key, payload)
					if err != nil {
						fmt.Printf("[worker %d] Put error: %v", workerID, err)
						return
					}
					val, err := s.frontend.Get(ctx, key)
					if err != nil {
						fmt.Printf("[worker %d] Get error: %v", workerID, err)
						return
					}
					if val != payload {
						fmt.Printf("[worker %d] Get returned unexpected value: %s != %s", workerID, val, payload)
						return
					}
					atomic.AddUint64(&reqCount, 2)
				}
			}
		}(id)
	}

	time.Sleep(dur)
	stop <- true
	close(stop)
	log.Println("Gonna wait")
	wg.Wait()
	log.Println("Finished all requests")
	return nil
}

func (s *workloadGen) ImplementsSimpleWorkload(context.Context) error {
	return nil
}
