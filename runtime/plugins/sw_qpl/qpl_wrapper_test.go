package sw_qpl

import (
	"bytes"
	crand "crypto/rand"
	rand "math/rand/v2"
	"testing"
)

// TestCompressDecompressRandom generates random data, compresses and decompresses it, and verifies integrity.
func TestCompressDecompressRandom(t *testing.T) {
	for i := range 1000 {
		// random data of random size up to 4096 bytes
		size := rand.IntN(4096)
		data := make([]byte, size)
		if _, err := crand.Read(data); err != nil {
			t.Fatalf("rand.Read failed: %v", err)
		}

		compressed, err := Compress(data)
		if err != nil {
			t.Fatalf("Compress failed on iteration %d: %v", i, err)
		}

		decompressed, err := Decompress(compressed, len(data))
		if err != nil {
			t.Fatalf("Decompress failed on iteration %d: %v", i, err)
		}

		if !bytes.Equal(data, decompressed) {
			t.Errorf("Data mismatch on iteration %d: original and decompressed differ", i)
			t.Errorf("Original: %x", data)
			t.Errorf("Decompressed: %x", decompressed)
		}
	}
}
