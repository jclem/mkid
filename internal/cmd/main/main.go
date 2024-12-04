// Package main runs the CLI.
package main

import (
	"context"
	"log"

	"github.com/jclem/mkid/internal/cmd"
)

func main() {
	ctx := context.Background()

	if err := cmd.Execute(ctx); err != nil {
		log.Fatalf("execute command: %s", err)
	}
}
