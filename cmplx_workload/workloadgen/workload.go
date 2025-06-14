package workloadgen

import (
	"context"
	"flag"
	"fmt"
	"golang.org/x/exp/rand"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/blueprint-uservices/blueprint/runtime/core/workload"

	"github.com/liam0215/anarres/workflow/frontend"
)

// Workload specific flags
var outfile = flag.String("outfile", "stats.csv", "Outfile where individual request information will be stored")
var dur = flag.String("dur", "15s", "Duration for which the workload should be run")
var tput = flag.Int64("tput", 4000, "Desired throughput")
var size = flag.Int("size", 64, "Size of value in KB")

type ComplexWorkload interface {
	ImplementsComplexWorkload(ctx context.Context) error
}

type complexWldGen struct {
	ComplexWorkload

	frontend frontend.Frontend
	value    string
}

func NewComplexWorkload(ctx context.Context, frontend frontend.Frontend) (ComplexWorkload, error) {
	w := &complexWldGen{frontend: frontend}
	return w, nil
}

type FnType func() error

func statWrapper(fn FnType) workload.Stat {
	start := time.Now()
	err := fn()
	duration := time.Since(start)
	s := workload.Stat{}
	s.Start = start.UnixNano()
	s.Duration = duration.Nanoseconds()
	s.IsError = (err != nil)
	return s
}

func (w *complexWldGen) RunPutHandler(ctx context.Context) workload.Stat {
	id := rand.Intn(500)
	key := "user" + strconv.Itoa(id)
	return statWrapper(func() error {
		return w.frontend.Put(ctx, key, w.value)
	})
}

func (w *complexWldGen) RunGetHandler(ctx context.Context) workload.Stat {
	id := rand.Intn(500)
	key := "user" + strconv.Itoa(id)
	return statWrapper(func() error {
		_, err := w.frontend.Get(ctx, key)
		return err
	})
}

func (w *complexWldGen) Run(ctx context.Context) error {
	f, err := os.Open("xml")
	if err != nil {
		return fmt.Errorf("failed to open xml file: %w", err)
	}

	limited := io.LimitReader(f, int64(*size)*1024)

	buf, err := io.ReadAll(limited)
	if err != nil {
		return fmt.Errorf("failed to read from xml: %w", err)
	}
	f.Close()

	w.value = string(buf)
	for i := range 500 {
		key := "user" + strconv.Itoa(i)
		if err := w.frontend.Put(ctx, key, w.value); err != nil {
			return err
		}
	}

	wrk := workload.NewWorkload()
	// Configure the workload with the client side generators for the various APIs and their respective proportions
	wrk.AddAPI("PutHandler", w.RunPutHandler, 50)
	wrk.AddAPI("GetHandler", w.RunGetHandler, 50)
	// Initialize the engine
	engine, err := workload.NewEngine(*outfile, *tput, *dur, wrk)
	if err != nil {
		return err
	}
	// Run the workload
	engine.RunOpenLoop(ctx)
	// Print statistics from the workload
	return engine.PrintStats()
}

func (w *complexWldGen) ImplementsComplexWorkload(context.Context) error {
	return nil
}
