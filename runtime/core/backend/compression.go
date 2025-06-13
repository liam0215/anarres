package backend

import "context"

type Compression interface {
	Compress(ctx context.Context, data []byte) ([]byte, error)

	Decompress(ctx context.Context, compressedData []byte, expectedLen int) ([]byte, error)
}
