package main

import (
	"github.com/htw-swa-jk-nk-ns/service-calculate/cmd"
	"github.com/rs/zerolog"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	cmd.Execute()
}
