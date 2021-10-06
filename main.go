package main

import (
	"context"
	"os"

	"slack-docker/cmd"
)

func main() {
	os.Exit(cmd.Run(context.Background(), os.Args))
}
