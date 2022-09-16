package queue

import (
	"github.com/MNThomson/securitybot/pkg/queue/sqs"
	AWSremoveSQS "github.com/aws/aws-sdk-go/service/sqs"
)

func Init() {
	sqs.InitSQS()
}

func PollMessages(chn chan<- *AWSremoveSQS.Message, MaxNumberOfMessages int64) {
	sqs.PollMessages(chn, MaxNumberOfMessages)
}

func DeleteMessage(ReceiptHandle string) error {
	sqs.DeleteMessage(ReceiptHandle)
	return nil
}

func SendMsg(message string) error {
	sqs.SendMsg(message)
	return nil
}
