//go:generate go tool mockery

package main

import (
	"log"

	"github.com/kafkaphoenix/gosurf/cmd/app/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
