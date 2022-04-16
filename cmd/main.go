package main

import (
	"codex/internal/app"
	"codex/internal/pkg/utils/config"
)

func main() {
	config.DevConfigStore.FromJson()
	config.ProdConfigStore.FromJson()

	app.RunServer()
}
