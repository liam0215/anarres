package frontend

import (
	"context"
	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"github.com/liam0215/anarres/workflow/compress"
	"github.com/pkg/errors"
)

type (
	Frontend interface {
		// Put a key-value pair in the cache.
		Put(ctx context.Context, key string, value string) error

		// Get a value by key from the cache.
		Get(ctx context.Context, key string) (string, error)
	}
)

type frontend struct {
	compress compress.CompressService
	kvCache  backend.Cache
}

// Instantiates the Frontend service, which makes calls to the compress service
func NewFrontend(ctx context.Context, compress compress.CompressService, cache backend.Cache) (Frontend, error) {
	f := &frontend{
		compress: compress,
		kvCache:  cache,
	}
	return f, nil
}

// Put implements Frontend.
func (f *frontend) Put(ctx context.Context, key string, value string) error {
	if key == "" {
		return errors.New("FrontendService.Put key cannot be empty")
	}

	compressed_value, err := f.compress.Compress(ctx, value)
	if err != nil {
		return err
	}

	err = f.kvCache.Put(ctx, key, compressed_value)
	return err
}

// Get implements Frontend.
func (f *frontend) Get(ctx context.Context, key string) (string, error) {
	if key == "" {
		return "", errors.New("FrontendService.Get key cannot be empty")
	}

	var compressed_value string
	got_value, err := f.kvCache.Get(ctx, key, &compressed_value)
	if err != nil {
		return "", err
	}
	if !got_value {
		return "", errors.New("FrontendService.Get key not found")
	}

	decompressed_value, err := f.compress.Decompress(ctx, compressed_value)
	if err != nil {
		return "", err
	}

	return decompressed_value, nil
}
