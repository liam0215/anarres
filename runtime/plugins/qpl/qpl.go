package qpl

import (
	"context"
	blueprintBackend "github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"github.com/liam0215/anarres/runtime/core/backend"
	"go.opentelemetry.io/otel/metric"
)

// Implements the backend.Compression interface
type QplCompression struct {
	backend.Compression

	meter             metric.Meter
	compressionSize   metric.Int64Histogram
	decompressionSize metric.Int64Histogram
}

// Instantiates a new qpl compression service
func NewQplCompression(ctx context.Context) (*QplCompression, error) {
	// Get OpenTelemetry meter
	meter, err := blueprintBackend.Meter(ctx, "compression")
	if err != nil {
		return nil, err
	}

	// Create metrics
	compressionSize, _ := meter.Int64Histogram("compression.compress_size_bytes")
	decompressionSize, _ := meter.Int64Histogram("compression.decompress_size_bytes")

	return &QplCompression{
		meter:             meter,
		compressionSize:   compressionSize,
		decompressionSize: decompressionSize,
	}, nil
}

// Implements the backend.Compression interface
func (q *QplCompression) Compress(ctx context.Context, data []byte) ([]byte, error) {
	compressed, err := Compress(data)
	if err != nil {
		return nil, err
	}

	// Record metrics
	originalLen := int64(len(data))
	q.compressionSize.Record(ctx, originalLen)

	return compressed, nil
}

// Implements the backend.Compression interface
func (q *QplCompression) Decompress(ctx context.Context, compressedData []byte, expectedLen int) ([]byte, error) {
	decompressed, err := Decompress(compressedData, expectedLen)
	if err != nil {
		return nil, err
	}

	// Record metrics
	decompressedLen := int64(len(decompressed))
	q.decompressionSize.Record(ctx, decompressedLen)

	return decompressed, nil
}
