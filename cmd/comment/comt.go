package main

import (
	mcscomt "codex/internal/app/comment"
	"codex/internal/pkg/utils/config"
	"codex/internal/pkg/utils/log"
)

func main() {
	err := config.DevConfigStore.FromJson()
	if err != nil {
		log.Error(err)
	}

	err = config.ProdConfigStore.FromJson()
	if err != nil {
		log.Error(err)
	}

	mcscomt.RunServer()
}
