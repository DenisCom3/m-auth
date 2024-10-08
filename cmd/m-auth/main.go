package main

import (
	"context"
	"github.com/DenisCom3/m-auth/internal/app"
	"log"
)

func main() {

	ctx := context.Background()

	a, err := app.New(ctx)

	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = a.Run(ctx)
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}

}
