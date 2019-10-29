package main

import (
	"context"
	"fmt"
	"os"

	"github.com/hlts2/create-go-app/pkg/create-go-app/cmd"
)

func main() {
	cmd := cmd.NewCommand(context.Background())
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "The application has terminated, because an error occurred: %v", err)
		os.Exit(1)
	}
}
