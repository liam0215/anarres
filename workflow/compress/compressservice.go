package compress

import (
	"context"
	"fmt"
	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"github.com/pkg/errors"
)

type (
	CompressService interface {
		// Compress value and return it
		Compress(ctx context.Context, value string) (string, error)

		// Decompress value and return it
		Decompress(ctx context.Context, value string) (string, error)
	}
)

type compressImpl struct {
	// TODO: Maybe add compress library instance here?
}

// Instantiates the Frontend service, which makes calls to the compress service
func NewCompressService(ctx context.Context) (CompressService, error) {
	c := &compressImpl{}
	return c, nil
}

// Compress implements CompressService.
func (c *compressImpl) Compress(ctx context.Context, value string) (string, error) {
	if value == "" {
		return errors.New("CompressService.Compress value cannot be empty")
	}

	// TODO: Implement calls to qpl wrapper

	return value, nil
}

// Decompress implements CompressService.
func (c *compressImpl) Decompress(ctx context.Context, value string) (string, error) {
	if value == "" {
		return errors.New("CompressService.Decompress value cannot be empty")
	}

	// TODO: Implement calls to qpl wrapper

	return value, nil
}
