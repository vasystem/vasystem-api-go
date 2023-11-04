package vasystemapi

import (
	"context"
	"crypto/tls"
	"fmt"

	vasystemapipb "buf.build/gen/go/vasystem/api/grpc/go/vasystem/api/v2/apiv2grpc"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

type Client struct {
	conn *grpc.ClientConn
}

func New(domain, clientId, clientSecret string, dialOpts ...grpc.DialOption) (*Client, error) {
	wellKnown, err := getWellKnown("https://" + domain)
	if err != nil {
		return nil, fmt.Errorf("failed to get well known: %w", err)
	}

	clientCredentialsConfig := &clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		AuthStyle:    oauth2.AuthStyleInHeader,
		Scopes:       []string{"airlines", "routes"},
		TokenURL:     wellKnown.TokenEndpoint,
	}

	dialOpts = append([]grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
		grpc.WithPerRPCCredentials(&oauth.TokenSource{
			TokenSource: clientCredentialsConfig.TokenSource(context.Background()),
		}),
	}, dialOpts...)

	conn, err := grpc.Dial(fmt.Sprintf("api.%s:443", domain), dialOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %w", err)
	}

	return &Client{
		conn: conn,
	}, nil
}

func (c *Client) Airlines() vasystemapipb.AirlinesServiceClient {
	return vasystemapipb.NewAirlinesServiceClient(c.conn)
}

func (c *Client) Routes() vasystemapipb.RoutesServiceClient {
	return vasystemapipb.NewRoutesServiceClient(c.conn)
}

func (c *Client) VirtualAirlines() vasystemapipb.VirtualAirlinesServiceClient {
	return vasystemapipb.NewVirtualAirlinesServiceClient(c.conn)
}
