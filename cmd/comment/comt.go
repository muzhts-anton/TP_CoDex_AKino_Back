package main

import (
	"codex/internal/app/comment"
	"codex/internal/pkg/utils/config"
)

func main() {
	config.DevConfigStore.FromJson()
	config.ProdConfigStore.FromJson()

	mcscomt.RunServer()
}
