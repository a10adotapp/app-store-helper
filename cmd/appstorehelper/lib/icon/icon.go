package icon

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/a10adotapp/appstorehelper/cmd/appstorehelper/lib/contexts"
	"github.com/a10adotapp/appstorehelper/cmd/appstorehelper/lib/openai"
	"github.com/a10adotapp/appstorehelper/cmd/appstorehelper/lib/openai/images"
	"github.com/a10adotapp/appstorehelper/cmd/appstorehelper/lib/ptr"
)

func GenerateIcon(ctx context.Context) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	promptBuf, err := os.ReadFile(filepath.Join(wd, "prompt.txt"))
	if err != nil {
		panic(err)
	}

	apiToken, err := contexts.APITokenFromContext(ctx)
	if err != nil {
		panic(err)
	}

	openAI, err := openai.NewOpenAI(ctx, &openai.NewOpenAIOptions{
		APIToken: apiToken,
	})
	if err != nil {
		panic(err)
	}

	err = openAI.Images.Generate(ctx, &images.GenerateOptions{
		Prompt: string(promptBuf),
		Model:  ptr.P("dall-e-2"),
	})

	fmt.Printf("err: %v\n", err)
}
