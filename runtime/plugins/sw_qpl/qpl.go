package sw_qpl

import (
	"context"
	"github.com/liam0215/anarres/runtime/core/backend"
)

// Implements the backend.Compression interface
type SwQplCompression struct {
	backend.Compression
}

// Instantiates a new qpl compression service
func NewSwQplCompression(ctx context.Context) (*SwQplCompression, error) {
	return &SwQplCompression{}, nil
}

// Implements the backend.Compression interface
func (q *SwQplCompression) Compress(ctx context.Context, data []byte) ([]byte, error) {
	compressed, err := Compress(data)
	if err != nil {
		return nil, err
	}

	return compressed, nil
}

// Implements the backend.Compression interface
func (q *SwQplCompression) Decompress(ctx context.Context, compressedData []byte, expectedLen int) ([]byte, error) {
	decompressed, err := Decompress(compressedData, expectedLen)
	if err != nil {
		return nil, err
	}

	return decompressed, nil
}
