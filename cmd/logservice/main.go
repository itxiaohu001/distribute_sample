package main

import (
	"context"
	"distribute_sample/log"
	"distribute_sample/registry"
	"distribute_sample/service"
	"fmt"
	stdlog "log"
)

func main() {
	log.Run("./service.log")

	host, port := "localhost", "8080"
	serviceUrl := fmt.Sprintf("http://%s:%s", host, port)
	reg := registry.Registration{
		ServiceName: "Log service",
		ServiceUrl:  serviceUrl,
	}

	ctx, err := service.Start(context.Background(), host, port, reg, log.RegisterHandler)
	if err != nil {
		stdlog.Fatal(err)
	}
	<-ctx.Done()

	fmt.Println("shutting done log service")
}
