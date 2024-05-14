package main

import (
	"context"
	"os"
)

func main() {
	component := Index()
	component.Render(context.Background(), os.Stdout)
}
