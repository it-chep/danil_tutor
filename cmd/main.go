package main

import (
	"context"

	"github.com/it-chep/danil_tutor.git/internal"
)

func main() {
	ctx := context.Background()

	internal.New(ctx).Run(ctx)
}
