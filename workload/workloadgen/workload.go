package workloadgen

import (
	"context"
	"flag"
	"fmt"
	"strconv"
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

var myarg = flag.Int("myarg", 12345, "help message for myarg")

func NewSimpleWorkload(ctx context.Context, frontend frontend.Frontend) (SimpleWorkload, error) {
	return &workloadGen{frontend: frontend}, nil
}

func (s *workloadGen) Run(ctx context.Context) error {
	fmt.Printf("myarg is %v\n", *myarg)
	value := 0
	err := s.frontend.Put(ctx, strconv.Itoa(*myarg), strconv.Itoa(value))
	if err != nil {
		return fmt.Errorf("error putting item: %w", err)
	}
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ctx.Done():
			return nil
		case t := <-ticker.C:
			fmt.Println("Tick at", t)
			val, err := s.frontend.Get(ctx, strconv.Itoa(*myarg))
			if err != nil {
				return err
			}
			fmt.Println("Key: ", *myarg, ", Value: ", val)
			value++
			err = s.frontend.Put(ctx, strconv.Itoa(*myarg), strconv.Itoa(value))
			if err != nil {
				return fmt.Errorf("error putting item: %w", err)
			}
		}
	}
}

func (s *workloadGen) ImplementsSimpleWorkload(context.Context) error {
	return nil
}
