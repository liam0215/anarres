package scheduler

import (
	"context"
	"time"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"github.com/liam0215/anarres/workflow/compress"
	"github.com/pkg/errors"
)

type (
	SchedulerService interface {
		Run(ctx context.Context) error
	}
)

type SchedulerServiceImpl struct {
	compress compress.CompressService
	logger   backend.Logger
}

func NewSchedulerServiceImpl(ctx context.Context, compress compress.CompressService) (*SchedulerServiceImpl, error) {
	logger := backend.GetLogger()
	return &SchedulerServiceImpl{
		compress: compress,
		logger:   logger,
	}, nil
}

// Run implements Scheduler.
func (s *SchedulerServiceImpl) Run(ctx context.Context) error {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			metrics, err := s.compress.GetMetrics(ctx)
			if err != nil {
				return errors.Wrap(err, "failed to get compression metrics")
			}
			if metrics.NumCompressions > 0 && metrics.CompressionSizeAcc > 0 {
				average_compression_size := metrics.CompressionSizeAcc / metrics.NumCompressions
				s.logger.Info(ctx, "Average compression size: %d", average_compression_size)
			} else if metrics.NumDecompressions > 0 && metrics.DecompressionSizeAcc > 0 {
				average_decompression_size := metrics.DecompressionSizeAcc / metrics.NumDecompressions
				s.logger.Info(ctx, "Average decompression size: %d", average_decompression_size)
			}
		}
	}
}
