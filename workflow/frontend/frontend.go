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

// Warning: fields must start with capital letters to actually be put in the cache due to json.Marshal semantics in Memcached plugin
type compressed_value struct {
	Comp      []byte
	DecompLen int
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

	if value == "" {
		return errors.New("FrontendService.Put value cannot be empty")
	}

	comp, err := f.compress.Compress(ctx, value)
	if err != nil {
		return err
	}
	if len(comp) == 0 {
		return errors.New("FrontendService.Put compressed value is empty")
	}

	compressed_value := compressed_value{
		Comp:      comp,
		DecompLen: len(value),
	}

	err = f.kvCache.Put(ctx, key, compressed_value)
	return err
}

// Get implements Frontend.
func (f *frontend) Get(ctx context.Context, key string) (string, error) {
	if key == "" {
		return "", errors.New("FrontendService.Get key cannot be empty")
	}

	var compressed_value compressed_value
	got_value, err := f.kvCache.Get(ctx, key, &compressed_value)
	if err != nil {
		return "", err
	}
	if !got_value {
		return "", errors.New("FrontendService.Get key not found")
	}
	if len(compressed_value.Comp) == 0 {
		return "", errors.New("FrontendService.Get Cache value is empty")
	}

	decomp, err := f.compress.Decompress(ctx, compressed_value.Comp, compressed_value.DecompLen)
	if err != nil {
		return "", err
	}

	return decomp, nil
}
