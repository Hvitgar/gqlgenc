module example

replace github.com/Hvitgar/gqlgenc => ../

go 1.15

require (
	github.com/99designs/gqlgen v0.16.0
	github.com/Hvitgar/gqlgenc v0.0.0-00010101000000-000000000000
	github.com/gorilla/websocket v1.4.2
	github.com/pkg/errors v0.9.1
	github.com/shopspring/decimal v1.4.0 // indirect
	github.com/stretchr/testify v1.4.0
	github.com/urfave/cli/v2 v2.3.0
	github.com/vektah/gqlparser/v2 v2.5.11
	golang.org/x/tools v0.22.0
	nhooyr.io/websocket v1.8.7
)
