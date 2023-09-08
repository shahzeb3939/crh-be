# export dynamodbTable=crh-lwrdag
GOOS=linux GOARCH=amd64 go build main.go
zip main.zip main
aws-vault exec DevAccount -- aws lambda update-function-code --function-name crh-be-lwrdag --zip-file fileb://main.zip