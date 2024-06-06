package main

import (
	"context"
	"fmt"
	"github.com/hasura/go-graphql-client"
	"github.com/johannessarpola/go-graphql-client-test/internal/app"
	"net/http"
)

type ApiKeyTransport struct {
	Transport http.RoundTripper
	Headers   map[string]string
}

func (c *ApiKeyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for key, value := range c.Headers {
		req.Header.Add(key, value)
	}
	return c.Transport.RoundTrip(req)
}

func main() {
	config, err := app.LoadConfig[app.Config]("config/config.dev.yaml")
	if err != nil {
		fmt.Println("could not load config", err)
		panic(err)
	}

	apt := &ApiKeyTransport{
		Transport: http.DefaultTransport,
		Headers: map[string]string{
			"X-API-Key": config.API.Key,
		},
	}

	httpClient := &http.Client{
		Transport: apt,
	}

	client := graphql.NewClient(config.API.Address, httpClient)

	var q struct {
		FeaturedPlaylists []struct {
			Id          string
			Name        string
			Description string
		}
	}
	err = client.Query(context.Background(), &q, nil)
	if err != nil {
		fmt.Println("could not do query", err)
		panic(err)
	}

	for _, featuredPlaylist := range q.FeaturedPlaylists {
		fmt.Printf("%s - %s\n %s\n", featuredPlaylist.Id, featuredPlaylist.Name, featuredPlaylist.Description)
	}

}
