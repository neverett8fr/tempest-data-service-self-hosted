package compression

import "net/http"

const (
	serviceSubPath    = "compression"
	compressSubPath   = "compress"
	decompressSubPath = "decompress"

	compressionGZip = "application/gzip"
	compressionXTar = "application/x-tar"
)

// check if compressed
func (cp *CompressionProvider) IsCompressed(data []byte) bool {

	supportedCompression := map[string]bool{
		compressionGZip: true,
		compressionXTar: true,
	}

	return supportedCompression[http.DetectContentType(data)]
}
