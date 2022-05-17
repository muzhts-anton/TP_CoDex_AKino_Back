package main

import (
	mcsrtng "codex/internal/app/rating"
	"codex/internal/pkg/utils/config"
)

func main() {
	config.DevConfigStore.FromJson()
	config.ProdConfigStore.FromJson()

	mcsrtng.RunServer()
}
