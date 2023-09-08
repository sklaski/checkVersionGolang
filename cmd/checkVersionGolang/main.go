package main

import (
	"context"
	"log"

	"checkVersionGolang/application"
)

func main() {
	ctx := context.Background()
	err := application.CompareVersions(ctx)
	if err != nil {
		log.Fatalf("an error occurred: %s", err)
	}
}
