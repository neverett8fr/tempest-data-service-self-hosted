package compression

import (
	"fmt"
)

// "convert" - convert either to compressed, or uncompressed
func (cp *CompressionProvider) Convert(data []byte) ([]byte, error) {

	if cp.IsCompressed(data) {
		return cp.Decompress(data)
	}

	return cp.Compress(data)
}

// call service to compress
func (cp *CompressionProvider) Compress(data []byte) ([]byte, error) {

	if cp.IsCompressed(data) {
		return nil, fmt.Errorf("err, data is already compressed")
	}

	return cp.callCompress(data)
}

// call service to decompress
func (cp *CompressionProvider) Decompress(data []byte) ([]byte, error) {

	if !cp.IsCompressed(data) {
		return nil, fmt.Errorf("err, data is already uncompressed")
	}

	return cp.callDecompress(data)
}
