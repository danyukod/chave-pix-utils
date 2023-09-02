package main

import (
	"context"
	"github.com/danyukod/chave-pix-utils/pkg/rabbitmq"
)

func main() {
	ctx := context.Background()

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publish(ctx, ch, []byte("Hello World!"))

}
