package main

//go:generate redocly bundle spec/main.yaml --output openapi.yaml
//go:generate go tool oapi-codegen -config oapi-codegen.yaml openapi.yaml
