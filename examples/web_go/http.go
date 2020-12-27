package main

import (
	"context"
	"log"

	"github.com/influx6/npkg/nhttp"
)

func servePage(ctx *nhttp.Ctx) error {

	return nil
}

func main() {
	var server = nhttp.NewServer(nhttp.HTTPFunc(servePage))

	var ctx = context.Background()
	if err := server.Listen(ctx, ":8060"); err != nil {
		log.Fatal(err)
	}
}
