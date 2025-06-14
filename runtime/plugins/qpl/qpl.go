package qpl

import (
	"context"
	"github.com/liam0215/anarres/runtime/core/backend"
)

// Implements the backend.Compression interface
type QplCompression struct {
	backend.Compression
}

// Instantiates a new qpl compression service
func NewQplCompression(ctx context.Context) (*QplCompression, error) {
	return &QplCompression{}, nil
}

// Implements the backend.Compression interface
func (q *QplCompression) Compress(ctx context.Context, data []byte) ([]byte, error) {
	compressed, err := Compress(data)
	if err != nil {
		return nil, err
	}

	return compressed, nil
}

// Implements the backend.Compression interface
func (q *QplCompression) Decompress(ctx context.Context, compressedData []byte, expectedLen int) ([]byte, error) {
	decompressed, err := Decompress(compressedData, expectedLen)
	if err != nil {
		return nil, err
	}

	return decompressed, nil
}
