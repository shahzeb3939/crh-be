package utils

import (
	"encoding/json"

	"github.com/shahzeb3939/crh-be/models"

	"github.com/aws/aws-lambda-go/events"
)

func ResponseObject(status int, msg string) (events.APIGatewayProxyResponse, error) {
	responseBody := models.ResponseBody{
		Message: msg,
	}

	responseBodyByteArray, err := json.Marshal(responseBody)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	headers := map[string]string{
		"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
		"Access-Control-Allow-Methods": "*",
		"Access-Control-Allow-Origin":  "*",
	}

	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Headers:    headers,
		Body:       string(responseBodyByteArray),
	}, nil
}
