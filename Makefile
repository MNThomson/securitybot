build:
	go build

events: *.go
	go run main.go events

ingestor: *.go
	go run main.go ingestor

## Dev Comamands
addQueueMessage:
	aws sqs send-message --queue-url "QUEUE URL" --message-body "This is a test message"

## Test Commands
