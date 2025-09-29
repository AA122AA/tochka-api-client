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

	opList, err := c.Acquiring.GetPayments(ctx, "1234567ab")
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Payment List - %+v\n", opList)
	log.Println()

	opInfo, err := c.Acquiring.GetOperationInfo(ctx, "48232c9a-ce82-1593-3cb6-5c85a1ffef8f")
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Operation info - %+v\n", opInfo)
	log.Println()

	registry, err := c.Acquiring.GetPaymentRegistry(ctx, "1234567ab", "200000000001097", "", "2025-09-23")
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Payment Registry - %+v\n", registry)
	log.Println()

	retailers, err := c.Acquiring.GetRetailers(ctx, "1234567ab")
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Retailers - %+v\n", retailers)
	log.Println()
}
