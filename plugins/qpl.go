package qpl

/*
#cgo LDFLAGS: -lqpl -lqpl_shim
#include <qpl/qpl.h>
#include <stdint.h>

int compress_software(uint8_t *src, uint8_t *dst, uint32_t src_size,
		      uint32_t dst_size);
uint32_t get_safe_compression_buffer_size(uint32_t src_size);
*/
import "C"
import (
	"unsafe"
)
import "context"


type CompressLibInterface interface {
	Compress(ctx context.Context, src []byte) ([]byte, error)

	Decompress(ctx context.Context, src []byte) ([]byte, error)
}

type QPL struct {
	CompressLib CompressLibInterface
}

func NewQPL(ctx context.Context) (*QPL, error) {
	qpl := &QPL{}
	return qpl, nil
}

func CompressLib(spec wiring.WiringSpec, name string) string {
	return define[CompressLibInterface, simplecache.SimpleCache](spec, name)
}

func (qpl *QPL) Compress(ctx context.Context, src []byte) ([]byte, error) {
	src_len := C.uint32_t(len(src))
	dst_len := C.get_safe_compression_buffer_size(src_len)
	dst := make([]byte, dst_len)

	ret := C.compress_software(
		(*C.uint8_t)(unsafe.Pointer(&src[0])),
		(*C.uint8_t)(unsafe.Pointer(&dst[0])),
		src_len,
		dst_len,
	)

	if ret != 0 {
		return nil, fmt.Errorf("compression failed")
	}
	return dst[:dst_len], nil
}

func define[CompressInterface any, CompressImpl any](spec wiring.WiringSpec, name string) string {
	compressName := name + ".compresslib"

	spec.Define(compressName, &{}, func(namespace wiring.Namespace) (ir.IRNode, error) {
		return newSimpleBackend[BackendImpl](name)
	})

	// Create a pointer to the backend instance
	pointer.CreatePointer[*SimpleBackend](spec, name, backendName)

	// Return the pointer; anybody who wants to access the backend instance should do so through the pointer
	return name
}
