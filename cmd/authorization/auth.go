package main

import (
	"codex/internal/app/authorization"
	"codex/internal/pkg/utils/config"
)

func main() {
	config.DevConfigStore.FromJson()
	config.ProdConfigStore.FromJson()

	mcsauth.RunServer()
}
