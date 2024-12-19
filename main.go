package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/99designs/gqlgen/api"
	"github.com/Hvitgar/gqlgenc/clientgen"
	"github.com/Hvitgar/gqlgenc/config"
	"github.com/Hvitgar/gqlgenc/generator"
)

func main() {
	ctx := context.Background()
	var cfg *config.Config
	var err error
	if path, ok := os.LookupEnv("GQLGEN_CONFIG"); ok {
		configFile, err := filepath.Abs(path)
		err = os.Chdir(filepath.Dir(configFile))
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to enter config dir: %w", err)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "absPath: %+v", err.Error())
			os.Exit(1)
		}

		cfg, err = config.LoadConfig(configFile)
	} else {
		cfg, err = config.LoadConfigFromDefaultLocations()
	}
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
