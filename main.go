//go:build lambda
// +build lambda

package main

import (
    "context"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"

    "nexus-enedis-m23-measure-lambda/internal/router"
)

var ginLambda *ginadapter.GinLambda

func init() {
    r := router.New()
    ginLambda = ginadapter.New(r)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
    lambda.Start(handler)
}

