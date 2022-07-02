package main

// #cgo CFLAGS: -g -Wall -Iinclude
// #cgo LDFLAGS: -L/home/viktor/Projects/Technopark/Backend/2022_1_CoDex/lib/linux -lvibesimple -lcurl -lssl -lvibecrypto -lfoo -lvibeictk -lvibeserver
// #include <stdio.h>
// #include <errno.h>

import (
	"codex/internal/app"
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

	app.RunServer()
}
