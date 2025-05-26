package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/liam0215/anarres/workflow/compress"
	"github.com/stretchr/testify/require"
)

// Tests acquire an CompressService instance using a service registry.
// This enables us to run local unit tests, while also enabling
// the Blueprint test plugin to auto-generate tests
// for different deployments when compiling an application.
var compressRegistry = registry.NewServiceRegistry[compress.CompressService]("compress_service")

func init() {
	// If the tests are run locally, we fall back to this CompressService implementation
	compressRegistry.Register("local", func(ctx context.Context) (compress.CompressService, error) {
		return compress.NewCompressService(ctx)
	})
}

func TestCompressService(t *testing.T) {
	ctx := context.Background()

	// Get the compress service
	compressService, err := compressRegistry.Get(ctx)

	// Try compressing a value
	val := "test_value"
	compressedVal, err := compressService.Compress(ctx, val)
	require.NoError(t, err)

	// Try decompressing the value
	decompressedVal, err := compressService.Decompress(ctx, compressedVal)
	require.NoError(t, err)

	require.Equal(t, decompressedVal, val)
}
