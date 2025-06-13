package compress

import (
	"context"
	"github.com/liam0215/anarres/runtime/core/backend"
	"github.com/pkg/errors"
)

type (
	CompressService interface {
		// Compress value and return it
		Compress(ctx context.Context, value string) ([]byte, error)

		// Decompress value and return it
		Decompress(ctx context.Context, value []byte, decompressedLen int) (string, error)
	}
)

type compressImpl struct {
	compressLib backend.Compression
}

// Instantiates the Frontend service, which makes calls to the compress service
func NewCompressService(ctx context.Context, compressLib backend.Compression) (CompressService, error) {
	c := &compressImpl{
		compressLib: compressLib,
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

	return string(dec), nil
}
