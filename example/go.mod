module example

replace github.com/Hvitgar/gqlgenc => ../

go 1.22.5

toolchain go1.23.4

require (
	github.com/99designs/gqlgen v0.16.0
	github.com/Hvitgar/gqlgenc v0.0.0-00010101000000-000000000000
	github.com/gorilla/websocket v1.5.0
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.10.0
	github.com/urfave/cli/v2 v2.27.5
	github.com/vektah/gqlparser/v2 v2.5.20
	golang.org/x/tools v0.28.0
	nhooyr.io/websocket v1.8.7
)

require (
	github.com/agnivade/levenshtein v1.2.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.5 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/klauspost/compress v1.10.3 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sosodev/duration v1.3.1 // indirect
	github.com/xrash/smetrics v0.0.0-20240521201337-686a1a2994c1 // indirect
	golang.org/x/mod v0.22.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// Replace github.com/99designs/gqlgen with local path
replace github.com/99designs/gqlgen => ../../gqlgen
