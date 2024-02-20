package openai

import (
	"context"

	"github.com/a10adotapp/appstorehelper/cmd/appstorehelper/lib/openai/config"
	"github.com/a10adotapp/appstorehelper/cmd/appstorehelper/lib/openai/images"
	"github.com/go-playground/validator/v10"
)

type OpenAI struct {
	Config *config.Config
	Images *images.Images
}

type NewOpenAIOptions struct {
	APIToken string `validate:"required"`
}

func NewOpenAI(ctx context.Context, opts *NewOpenAIOptions) (*OpenAI, error) {
	options, err := parseOptions(opts)
	if err != nil {
		panic(err)
	}

	config := &config.Config{
		APIToken: options.APIToken,
	}

	return &OpenAI{
		Config: config,
		Images: images.NewImages(config),
	}, nil
}

func parseOptions(opts *NewOpenAIOptions) (NewOpenAIOptions, error) {
	options := NewOpenAIOptions{}

	if opts != nil {
		options.APIToken = opts.APIToken
	}

	if err := validator.New().Struct(options); err != nil {
		return options, err
	}

	return options, nil
}
