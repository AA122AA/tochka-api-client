package main

import (
	"context"
	"log"

	"github.com/AA122AA/tochka-api-client/tochka"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := tochka.NewClient("https://enter.tochka.com/sandbox/v2", "v1.0", "sandbox.jwt.token")
	data, err := c.Acquiring.GetPayments(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(data)
}
