package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type GoServerlessBoilerplateStackProps struct {
	awscdk.StackProps
}

func NewGoServerlessBoilerplateStack(scope constructs.Construct, id string, props *GoServerlessBoilerplateStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	usersTable := awsdynamodb.NewTable(
		stack, 
		jsii.String("usersTable"), 
		&awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
				Name: jsii.String("email"),
				Type: awsdynamodb.AttributeType_STRING,
			},
			TableName: jsii.String("usersTable"),
	})

	userRegistrationHandlerFunction := awslambda.NewFunction(
		stack, 
		jsii.String("userRegistrationHandlerFunction"), 
		&awslambda.FunctionProps{
			FunctionName: jsii.String("userRegistrationHandlerFunction"),
			Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
			Code:         awslambda.AssetCode_FromAsset(jsii.String("src/."), nil),
			Handler:      jsii.String("main"),
		})

	usersTable.GrantReadWriteData(userRegistrationHandlerFunction)

	// API Gateway
	api := awsapigateway.NewRestApi(stack, jsii.String("restApi"), &awsapigateway.RestApiProps{
		DefaultCorsPreflightOptions: &awsapigateway.CorsOptions{
			AllowHeaders: jsii.Strings("Content-Type", "X-Amz-Date", "Authorization", "X-Api-Key"),
			AllowMethods: jsii.Strings("GET", "POST", "PUT", "DELETE", "OPTIONS"),
			AllowOrigins: jsii.Strings("*"),
		},
		DeployOptions: &awsapigateway.StageOptions{
			LoggingLevel:     awsapigateway.MethodLoggingLevel_INFO,
			DataTraceEnabled: jsii.Bool(true),
		},
		EndpointConfiguration: &awsapigateway.EndpointConfiguration{
			Types: &[]awsapigateway.EndpointType{awsapigateway.EndpointType_REGIONAL},
		},
	})

	registerRoute("register", "POST", &userRegistrationHandlerFunction, api)

	// userRegistrationFuncIntegration := awsapigateway.NewLambdaIntegration(userRegistrationHandlerFunction, nil)

	// '/register'
	// registerUserResource := api.Root().AddResource(jsii.String("register"), nil)
	// registerUserResource.AddMethod(jsii.String("POST"), userRegistrationFuncIntegration, nil)

	// '/login'
	// loginResource := api.Root().AddResource(jsii.String("login"), nil)
	// loginResource.AddMethod(jsii.String("POST"), integration, nil)

	return stack
}

func registerRoute(path, method string, handlerFunction *awslambda.Function, api awsapigateway.RestApi) {
	userRegistrationFuncIntegration := awsapigateway.NewLambdaIntegration(*handlerFunction, nil)

	registerUserResource := api.Root().AddResource(jsii.String(path), nil)
	registerUserResource.AddMethod(jsii.String(method), userRegistrationFuncIntegration, nil)
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewGoServerlessBoilerplateStack(app, "GoServerlessBoilerplateStack", &GoServerlessBoilerplateStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
