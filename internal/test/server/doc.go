package server

//go:generate go run github.com/andrewstucki/oapi-codegen/cmd/oapi-codegen --config=config.yaml ../test-schema.yaml
//go:generate go run github.com/matryer/moq -out server_moq.gen.go . ServerInterface
