package main

import (
	"context"
	"distribute_sample/log"
	"distribute_sample/service"
	"fmt"
	stdlog "log"
)

func main() {
	log.Run("./service.log")

	host, port := "localhost", "8080"

	ctx, err := service.Start(context.Background(), "log service", host, port, log.RegisterHandler)
	if err != nil {
		stdlog.Fatal(err)
	}
	<-ctx.Done()

	fmt.Println("shutting done log service")
}
