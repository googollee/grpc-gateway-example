package main

import (
	"context"
	"fmt"
	"log"

	"grpc_gateway/http_client/client"
	"grpc_gateway/http_client/client/service"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func main() {
	ctx := context.Background()

	transport := httptransport.New("localhost:8080", "", nil)

	// create the API client, with the transport
	client := client.New(transport, strfmt.Default)
	basicAuth := httptransport.BasicAuth("user", "password")

	name := "world"
	ticker := "1"
	arg := service.NewServiceRepeatGreetParamsWithContext(ctx).WithName(&name).WithTickerSecond(&ticker)
	resp, err := client.Service.ServiceRepeatGreet(arg, basicAuth)
	if err != nil {
		log.Fatal(err)
	}
	greet := resp.GetPayload()
	fmt.Println(greet.Error, greet.Result)
}
