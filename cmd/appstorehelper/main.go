package main

import (
	"context"
	"flag"
	"os"

	"github.com/a10adotapp/appstorehelper/cmd/appstorehelper/lib/contexts"
	"github.com/a10adotapp/appstorehelper/cmd/appstorehelper/lib/icon"
)

func main() {
	flag.Parse()

	ctx := context.Background()

	apiToken := os.Getenv("API_TOKEN")
	if apiToken == "" {
		panic("missing required environment value: API_TOKEN")
	}

	ctx = contexts.ContextWithAPIToken(ctx, apiToken)

	switch flag.Arg(0) {
	case "icongenerate":
		icon.GenerateIcon(ctx)
	}
}
