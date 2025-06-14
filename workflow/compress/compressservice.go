package compress

import (
	"context"
	"github.com/liam0215/anarres/runtime/core/backend"
	"github.com/pkg/errors"
	"sync/atomic"
)

type (
	CompressService interface {
		// Compress value and return it
		Compress(ctx context.Context, value string) ([]byte, error)

		// Decompress value and return it
		Decompress(ctx context.Context, value []byte, decompressedLen int) (string, error)

		// GetMetrics returns the compression and decompression metrics
		GetMetrics(ctx context.Context) (CompressionMetrics, error)
	}
)

type CompressionMetrics struct {
	CompressionSizeAcc   int64
	NumCompressions      int64
	DecompressionSizeAcc int64
	NumDecompressions    int64
}

type compressImpl struct {
	compressLib backend.Compression

	metrics CompressionMetrics
}

// Instantiates the Frontend service, which makes calls to the compress service
func NewCompressService(ctx context.Context, compressLib backend.Compression) (CompressService, error) {
	c := &compressImpl{
		compressLib: compressLib,
		metrics: CompressionMetrics{
			CompressionSizeAcc:   0,
			NumCompressions:      0,
			DecompressionSizeAcc: 0,
			NumDecompressions:    0,
		},
	}
	return c, nil
}

// Compress implements CompressService.
func (c *compressImpl) Compress(ctx context.Context, value string) ([]byte, error) {
	if value == "" {
		return []byte(""), errors.New("CompressService.Compress value cannot be empty")
	}

	comp, err := c.compressLib.Compress(ctx, []byte(value))
	if err != nil {
		panic(err)
	}

	// Record metrics
	originalLen := int64(len(value))
	atomic.AddInt64(&c.metrics.CompressionSizeAcc, originalLen)
	atomic.AddInt64(&c.metrics.NumCompressions, 1)

	return comp, nil
}

// Decompress implements CompressService.
func (c *compressImpl) Decompress(ctx context.Context, value []byte, decompressedLen int) (string, error) {
	if len(value) == 0 {
		return "", errors.New("CompressService.Decompress value cannot be empty")
	}

	dec, err := c.compressLib.Decompress(ctx, value, decompressedLen)
	if err != nil {
		panic(err)
	}

	// Record metrics
	atomic.AddInt64(&c.metrics.DecompressionSizeAcc, int64(decompressedLen))
	atomic.AddInt64(&c.metrics.NumDecompressions, 1)

	return string(dec), nil
}

func (c *compressImpl) GetMetrics(ctx context.Context) (CompressionMetrics, error) {
	return c.metrics, nil
}
