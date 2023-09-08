# export dynamodbTable=crh-lwrdag
GOOS=linux GOARCH=amd64 go build main.go
zip main.zip main
aws-vault exec DevAccount -- aws s3 cp main.zip s3://crh-be-build/main.zip