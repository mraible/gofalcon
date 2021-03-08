package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/crowdstrike/gofalcon/falcon"
	"github.com/crowdstrike/gofalcon/falcon/client/event_streams"
	"github.com/crowdstrike/gofalcon/pkg/falcon_util"
)

func main() {
	clientId := flag.String("client-id", os.Getenv("FALCON_CLIENT_ID"), "Client ID for accessing CrowdStrike Falcon Platform (default taken from FALCON_CLIENT_ID env)")
	clientSecret := flag.String("client-secret", os.Getenv("FALCON_CLIENT_SECRET"), "Client Secret for accessing CrowdStrike Falcon Platform (default taken from FALCON_CLIENT_SECRET)")
	clientCloud := flag.String("cloud", os.Getenv("FALCON_CLOUD"), "Falcon cloud abbreviation (us-1, us-2, eu-1, us-gov-1)")

	flag.Parse()

	if *clientId == "" {
		*clientId = promptUser(`Missing FALCON_CLIENT_ID environment variable. Please provide your OAuth2 API Client ID for authentication with CrowdStrike Falcon platform. Establishing and retrieving OAuth2 API credentials can be performed at https://falcon.crowdstrike.com/support/api-clients-and-keys.
Falcon Client ID`)
	}
	if *clientSecret == "" {
		*clientSecret = promptUser(`Missing FALCON_CLIENT_SECRET environment variable. Please provide your OAuth2 API Client Secret for authentication with CrowdStrike Falcon platform. Establishing and retrieving OAuth2 API credentials can be performed at https://falcon.crowdstrike.com/support/api-clients-and-keys.
Falcon Client Secret`)
	}

	client, err := falcon.NewClient(&falcon.ApiConfig{
		ClientId:     *clientId,
		ClientSecret: *clientSecret,
		Cloud:        falcon.Cloud(*clientCloud),
		Context:      context.Background(),
	})
	if err != nil {
		panic(err)
	}

	json := "json"
	appId := "falcon_event_stream"
	// Step 1: Discover Available Streams
	response, err := client.EventStreams.ListAvailableStreamsOAuth2(&event_streams.ListAvailableStreamsOAuth2Params{
		AppID:   appId,
		Format:  &json,
		Context: context.Background(),
	})
	if err != nil {
		panic(falcon.ErrorExplain(err))
	}
	for _, e := range response.Payload.Errors {
		fmt.Println(e)
	}

	availableStreams := response.Payload.Resources
	for _, availableStream := range availableStreams {
		stream, err := falcon.NewStream(context.Background(), client, appId, availableStream)
		if err != nil {
			panic(err)
		}
		defer stream.Close()

		var fatal error
		for fatal == nil {
			select {
			case err := <-stream.Errors:
				if err.Fatal {
					fatal = err.Err
				} else {
					fmt.Fprintln(os.Stderr, err)
				}
			case event := <-stream.Events:
				pretty, err := falcon_util.PrettyJson(event)
				if err != nil {
					panic(err)
				}
				fmt.Println(pretty)
			}
		}
		panic(fatal)
	}
}

func promptUser(prompt string) string {
	fmt.Printf("%s: ", prompt)
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(s)
}
