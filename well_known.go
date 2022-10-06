package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var wellKnownEndpoint = &url.URL{
	Path: ".well-known/openid-configuration",
}

func getWellKnown(issuer string) (*WellKnownResponse, error) {
	baseAddr, err := url.Parse(issuer)
	if err != nil {
		return nil, fmt.Errorf("failed to parse server address: %w", err)
	}

	wellKnownUrl := baseAddr.ResolveReference(wellKnownEndpoint)

	res, err := http.Get(wellKnownUrl.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get configuration: %w", err)
	}
	defer res.Body.Close()

	var resp WellKnownResponse

	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, fmt.Errorf("failed to decode configuration: %w", err)
	}

	return &resp, nil
}

type WellKnownResponse struct {
	Issuer                string `json:"issuer"`
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
	RevocationEndpoint    string `json:"revocation_endpoint"`
	UserinfoEndpoint      string `json:"userinfo_endpoint"`
}
