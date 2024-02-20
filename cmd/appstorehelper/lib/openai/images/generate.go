package images

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"

	"github.com/a10adotapp/appstorehelper/cmd/appstorehelper/lib/contexts"
	"github.com/a10adotapp/appstorehelper/cmd/appstorehelper/lib/openai/config"
	"github.com/a10adotapp/appstorehelper/cmd/appstorehelper/lib/ptr"
	"github.com/go-playground/validator/v10"
)

type GenerateOptions struct {
	Prompt         string  `json:"prompt" validate:"required,max=1000"`
	Model          *string `json:"model,omitempty" validate:"omitempty,oneof=dall-e-2 dall-e-3"`
	N              *int    `json:"n,omitempty" validate:"omitempty,min=1,max=10"`
	Quality        *string `json:"quality,omitempty" validate:"omitempty,oneof=standard hd"`
	ResponseFormat *string `json:"response_format,omitempty" validate:"omitempty,oneof=url b64_json"`
	Size           *string `json:"size,omitempty" validate:"omitempty,oneof=256x256 512x512 1024x1024 1792x1024 1024x1792"`
	Style          *string `json:"style,omitempty" validate:"omitempty,oneof=vivid natural"`
	User           *string `json:"user,omitempty"`
}

func (i *Images) Generate(ctx context.Context, opts *GenerateOptions) error {
	options, err := parseOptions(opts)
	if err != nil {
		return err
	}

	jsonBuf, err := json.Marshal(options)
	if err != nil {
		return err
	}

	fmt.Printf("request: %s\n", string(jsonBuf))

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s", config.BaseURL, path.Join(BasePath, "generations")), bytes.NewReader(jsonBuf))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	apiToken, err := contexts.APITokenFromContext(ctx)
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	bodyBuf, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	fmt.Printf("response: %s\n", bodyBuf)

	return nil
}

func parseOptions(opts *GenerateOptions) (GenerateOptions, error) {
	options := GenerateOptions{
		Model: ptr.P("dall-e-3"),
	}

	fmt.Printf("opts: %+v\n", opts)

	if opts != nil {
		options.Prompt = opts.Prompt

		if opts.Model != nil {
			options.Model = opts.Model
		}

		if opts.N != nil {
			options.N = opts.N
		}

		if opts.Quality != nil {
			options.Quality = opts.Quality
		}

		if opts.ResponseFormat != nil {
			options.ResponseFormat = opts.ResponseFormat
		}

		if opts.Size != nil {
			options.Size = opts.Size
		}

		if opts.Style != nil {
			options.Style = opts.Style
		}

		if opts.User != nil {
			options.User = opts.User
		}
	}

	if err := validator.New().Struct(options); err != nil {
		return options, err
	}

	return options, nil
}
