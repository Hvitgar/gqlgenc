package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/api"
	"github.com/Hvitgar/gqlgenc/clientgen"
	"github.com/Hvitgar/gqlgenc/config"
	"github.com/Hvitgar/gqlgenc/generator"
	"os"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintf(os.Stderr, "cfg: %+v", err.Error())
		os.Exit(2)
	}

	clientGen := api.AddPlugin(clientgen.New(cfg))

	if err := generator.Generate(ctx, cfg, clientGen); err != nil {
		fmt.Fprintf(os.Stderr, "generate: %+v", err.Error())
		os.Exit(4)
	}
}
