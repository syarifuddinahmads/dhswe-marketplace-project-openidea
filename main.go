package main

import (
	"context"
	"log"

	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/http"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatalf("%+v", err)
	}
}

func run(ctx context.Context) error {
	server, err := http.NewServer()
	if err != nil {
		return err
	}
	err = server.Run(ctx)
	return err
}
