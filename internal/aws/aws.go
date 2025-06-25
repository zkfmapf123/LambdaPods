package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

type ClientParams struct {
	lambdaClient lambda.Client
}

func NewClient() ClientParams {

	client, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic(err)
	}

	return ClientParams{
		lambdaClient: *lambda.NewFromConfig(client),
	}
}
