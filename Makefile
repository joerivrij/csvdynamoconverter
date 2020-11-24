build:
	GO111MODULE=on GOOS=linux CGO_ENABLED=0 go build -o csvdynamoconverterctl cmd/main.go
	sudo mv csvdynamoconverterctl /usr/local/bin/csvdynamoconverterctl