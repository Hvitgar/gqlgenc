package clientgen

import (
	"fmt"

	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/codegen/templates"
)

func RenderTemplate(cfg *config.Config, types []*Type, ptrTypes []PtrType, operations []*Operation, generateClient bool, generateInterface bool, client config.PackageConfig) error {
	if err := templates.Render(templates.Options{
		PackageName: client.Package,
		Filename:    client.Filename,
		Data: map[string]interface{}{
			"Types":             types,
			"PtrTypes":          ptrTypes,
			"Operations":        operations,
			"GenerateClient":    generateClient,
			"GenerateInterface": generateInterface,
		},
		Packages:   cfg.Packages,
		PackageDoc: "// Code generated by github.com/Hvitgar/gqlgenc, DO NOT EDIT.\n",
	}); err != nil {
		return fmt.Errorf("%s generating failed: %w", client.Filename, err)
	}

	return nil
}
