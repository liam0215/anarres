package compress

/*
#cgo CPPFLAGS: -I/opt/qpl/include
#cgo LDFLAGS: -L/opt/qpl/lib -lqpl
#include "qpl_wrapper.h"
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

// Compress calls the QPL compression wrapper function in C.
// Note: the buffer must not contain any Go pointers, as it is passed directly to C.
func Compress(buffer []byte) ([]byte, error) {
	if len(buffer) == 0 {
		return nil, nil
	}

	var cOut *C.uint8_t
	var cOutLen C.uint32_t

	// use Go pointer directly to avoid copying
	status := C.qpl_compress_wrapper(
		(*C.uint8_t)(unsafe.Pointer(&buffer[0])),
		C.uint32_t(len(buffer)),
		&cOut,
		&cOutLen,
		C.qpl_path_auto,
	)
	if status != C.QPL_STS_OK {
		return nil, fmt.Errorf("compression failed: %d", status)
	}

	// transform to Go pointer
	out := C.GoBytes(unsafe.Pointer(cOut), C.int(cOutLen))
	C.free(unsafe.Pointer(cOut))
	return out, nil
}

// Decompress calls the QPL decompression wrapper function in C.
// expectedLen is the expected length of the decompressed data, which is necessary to allocate the output buffer correctly. It can be overestimated, but not underestimated.
func Decompress(buffer []byte, expectedLen int) ([]byte, error) {
	if len(buffer) == 0 {
		return nil, nil
	}

	var cOut *C.uint8_t
	var cOutLen C.uint32_t

	status := C.qpl_decompress_wrapper(
		(*C.uint8_t)(unsafe.Pointer(&buffer[0])),
		C.uint32_t(len(buffer)),
		C.uint32_t(expectedLen),
		&cOut,
		&cOutLen,
		C.qpl_path_auto,
	)
	if status != C.QPL_STS_OK {
		return nil, fmt.Errorf("decompression failed: %d", status)
	}

	// transform to Go pointer
	out := C.GoBytes(unsafe.Pointer(cOut), C.int(cOutLen))
	C.free(unsafe.Pointer(cOut))
	return out, nil
}
