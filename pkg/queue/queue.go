package queue

import (
	"github.com/MNThomson/securitybot/pkg/queue/sqs"
	"github.com/MNThomson/securitybot/pkg/types"
)

func Init() {
	sqs.InitSQS()
}

func PollMessages(chn chan<- types.AlertType, MaxNumberOfMessages int64) {
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
