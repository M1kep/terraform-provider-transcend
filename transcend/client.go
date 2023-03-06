package transcend

import (
	"net/http"

	graphql "github.com/hasura/go-graphql-client"
)

type backendTransport struct {
	apiToken    string
	internalKey string
}

func (t *backendTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+t.apiToken)
	if t.internalKey != "" {
		req.Header.Add("x-sombra-authorization", "Bearer "+t.internalKey)
	}
	return http.DefaultTransport.RoundTrip(req)
}

type Client struct {
	graphql      *graphql.Client
	sombraClient *http.Client
	url          string
}

func NewClient(url, apiToken string, internalKey string) *Client {
	client := &http.Client{Transport: &backendTransport{apiToken: apiToken, internalKey: internalKey}}

	return &Client{
		graphql:      graphql.NewClient(url, client),
		sombraClient: client,
		url:          url,
	}
}
