package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/protobuf/encoding/protojson"
	vasystemapi "vasystem-api-example-go/genproto/vasystem/api/v2"
)

func main() {
	domainPtr := flag.String("domain", "staralliancevirtual.org", "the main VA domain")
	clientIdPtr := flag.String("client-id", "client-id", "the client id")
	clientSecretPtr := flag.String("client-secret", "client-secret", "the client secret")

	flag.Parse()

	domain := *domainPtr

	wellKnown, err := getWellKnown("https://" + domain)
	if err != nil {
		log.Fatal(err)
	}

	clientCredentialsConfig := &clientcredentials.Config{
		ClientID:     *clientIdPtr,
		ClientSecret: *clientSecretPtr,
		AuthStyle:    oauth2.AuthStyleInHeader,
		Scopes:       []string{"airlines", "routes"},
		TokenURL:     wellKnown.TokenEndpoint,
	}

	conn, err := grpc.Dial(
		fmt.Sprintf("api.%s:443", domain),
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
		grpc.WithPerRPCCredentials(&oauth.TokenSource{
			TokenSource: clientCredentialsConfig.TokenSource(context.Background()),
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	airlinesClient := vasystemapi.NewAirlinesServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	airlines, err := airlinesClient.ListAirlines(ctx, &vasystemapi.ListAirlinesRequest{})
	if err != nil {
		log.Fatal(err)
	}

	data, err := protojson.MarshalOptions{
		Indent: "  ",
	}.Marshal(airlines)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
