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

var (
	TABLE_NAME   = "lambdapods"
	PRIMARY_KEY  = "function_arn"
	BILLING_MODE = true
)

// dynamoDB에 저장될 ARN 리스트
type FunctionAttrParams struct {
	FunctionArn string `json:"function_arn"`
	IAMArn      string `json:"iam_arn"`
	SQSArn      string `json:"sqs_arn"`
	SNSArn      string `json:"sns_arn"`
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

	// db client
	dynamoClient orm.TableParmas
}

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
