package main

import (
	"context"
	"distribute_sample/registry"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/services", &registry.RegistryService{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var svr = http.Server{}
	svr.Addr = registry.ServerPort

	go func() {
		log.Println(svr.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Printf("Registry service started, press any key to stop\n")
		var s string
		_, _ = fmt.Scanln(&s)
		_ = svr.Shutdown(ctx)
		cancel()
	}()

	<-ctx.Done()
	fmt.Printf("Shutting down registry service")
}
