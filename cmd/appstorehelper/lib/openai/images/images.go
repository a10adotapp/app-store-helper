package images

import (
	"github.com/a10adotapp/appstorehelper/cmd/appstorehelper/lib/openai/config"
)

const (
	BasePath = "images"
)

type Images struct {
	Config *config.Config
}

func NewImages(config *config.Config) *Images {
	return &Images{
		Config: config,
	}
}
