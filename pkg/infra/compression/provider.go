package compression

import (
	"context"
	"fmt"
	"tempest-data-service/pkg/config"
)

type CompressionProvider struct {
	Path           string
	UseCompression bool
}

func InitialiseCompressionProvider(ctx context.Context, conf config.Compression) (CompressionProvider, error) {

	pathToService := fmt.Sprintf("%s:%v/%s", conf.Host, conf.Port, serviceSubPath)
	if conf.Port == 0 {
		pathToService = fmt.Sprintf("%s/%s", conf.Host, serviceSubPath)
	}

	return CompressionProvider{
		Path: pathToService,
	}, nil
}
