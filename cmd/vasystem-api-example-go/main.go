package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	vasystemapipb "github.com/vasystem/api/grpc/go/vasystem/api/v2"
	"github.com/vasystem/vasystem-api-go/vasystemapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var marshaler = protojson.MarshalOptions{
	Indent: "  ",
}

func printMessage(msg proto.Message) {
	data, err := marshaler.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}

func main() {
	domainPtr := flag.String("domain", "staralliancevirtual.org", "the main VA domain")
	clientIdPtr := flag.String("client-id", "client-id", "the client id")
	clientSecretPtr := flag.String("client-secret", "client-secret", "the client secret")

	flag.Parse()

	client, err := vasystemapi.New(*domainPtr, *clientIdPtr, *clientSecretPtr)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	virtualAirline, err := client.VirtualAirlines().GetVirtualAirline(ctx, &vasystemapipb.GetVirtualAirlineRequest{})
	if err != nil {
		log.Fatal(err)
	}

	printMessage(virtualAirline)

	airlines, err := client.Airlines().ListAirlines(ctx, &vasystemapipb.ListAirlinesRequest{})
	if err != nil {
		log.Fatal(err)
	}

	printMessage(airlines)

	routes, err := client.Routes().ListRoutes(ctx, &vasystemapipb.ListRoutesRequest{})
	if err != nil {
		log.Fatal(err)
	}

	printMessage(routes)
}
