package compression

import (
	"encoding/binary"
	"log"
)

const (
	serviceSubPath    = "compression"
	compressSubPath   = "compress"
	decompressSubPath = "decompress"

	compressionGZip = "application/gzip"
	compressionXTar = "application/x-tar"
)

// check if compressed
// func (cp *CompressionProvider) IsCompressed(data []byte) bool {

// 	supportedCompression := map[string]bool{
// 		compressionGZip: true,
// 		compressionXTar: true,
// 	}

// 	return supportedCompression[http.DetectContentType(data)]
// }

func (cp *CompressionProvider) IsCompressed(data []byte) bool {
	if len(data) < 3 {
		log.Println("<3")
		return false
	}

	// Check for RLE compression
	if data[0] == 0x01 && data[1] == 0x00 {
		log.Println("rle")
		return true // ".rle"
	}

	// Check for gzip compression
	if data[0] == 0x1f && data[1] == 0x8b {
		log.Println("gzip")
		return true // ".gzip"
	}

	// Check for LZW compression
	if len(data) >= 4 && binary.LittleEndian.Uint32(data[0:4]) == 0x4C5A5700 {
		log.Println("lzw")
		return true
	}
	log.Println("false")
	return false
}
