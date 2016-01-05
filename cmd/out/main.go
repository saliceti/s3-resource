package main

import (
	"encoding/json"
	"os"

	"github.com/concourse/s3-resource"
	"github.com/concourse/s3-resource/out"
)

func main() {
	if len(os.Args) < 2 {
		s3resource.Sayf("usage: %s <sources directory>\n", os.Args[0])
		os.Exit(1)
	}

	var request out.OutRequest
	inputRequest(&request)

	sourceDir := os.Args[1]

	client, err := s3resource.NewS3Client(
		os.Stderr,
		request.Source.AccessKeyID,
		request.Source.SecretAccessKey,
		request.Source.RegionName,
		request.Source.Endpoint,
		request.Source.CredentialsSource,
	)
	if err != nil {
		s3resource.Fatal("building S3 client", err)
	}

	command := out.NewOutCommand(client)
	response, err := command.Run(sourceDir, request)
	if err != nil {
		s3resource.Fatal("running command", err)
	}

	outputResponse(response)
}

func inputRequest(request *out.OutRequest) {
	if err := json.NewDecoder(os.Stdin).Decode(request); err != nil {
		s3resource.Fatal("reading request from stdin", err)
	}
}

func outputResponse(response out.OutResponse) {
	if err := json.NewEncoder(os.Stdout).Encode(response); err != nil {
		s3resource.Fatal("writing response to stdout", err)
	}
}
