package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	orm "github.com/zkfmapf123/dynamoGORM"
)

type AWSParmas interface {
	Create(ClientParams)
	Update(ClientParams)
	Delete(ClientParams)
	Retrieve(ClientParams)
}

type ClientParams struct {
	// albClient
	// apiGWClient
	// mskClient
	// kdsClient
	// cronClient
	sqsClient    sqs.Client
	snsClient    sns.Client
	iamClient    iam.Client
	lambdaClient lambda.Client
	dynamoClient orm.TableParmas
}

var (
	TABLE_NAME   = "lambdapods"
	PRIMARY_KEY  = "function_arn"
	BILLING_MODE = true
)

func NewClient() ClientParams {
	c := context.Background()

	// aws
	client, err := config.LoadDefaultConfig(c)
	if err != nil {
		panic(err)
	}

	// dynamorm
	lambdaInfo := orm.TableParmas{
		TableName:   TABLE_NAME,
		Primarykey:  PRIMARY_KEY,
		BillingMode: BILLING_MODE,
	}

	return ClientParams{
		lambdaClient: *lambda.NewFromConfig(client),
		iamClient:    *iam.NewFromConfig(client),
		dynamoClient: lambdaInfo,
	}
}

type ServiceType string

const (
	LambdaType ServiceType = "lambda"
	SQSType    ServiceType = "sqs"
	SNSType    ServiceType = "sns"
	IAMType    ServiceType = "iam"
)

func CreateService(service ServiceType, c ClientParams) AWSParmas {
	switch service {
	case LambdaType:
		return NewLambda(c)

	case IAMType:
		return NewIAM(c)

	case SQSType:
		return NewSQS(c)

	case SNSType:
		return NewSns(c)

	default:
		return nil
	}

}
