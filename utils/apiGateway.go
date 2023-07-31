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

	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(responseBodyByteArray),
	}, nil
}
