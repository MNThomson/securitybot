package ingestor

import (
	"fmt"

	"github.com/MNThomson/securitybot/pkg/queue"
	AWSremoveSQS "github.com/aws/aws-sdk-go/service/sqs"
	log "github.com/sirupsen/logrus"
)

const (
	MaxNumberOfMessages int64 = 10
)

func handleMessage(msg *AWSremoveSQS.Message) {
	fmt.Println("RECEIVING MESSAGE >>> ")
	fmt.Println(*msg.Body)
}

func Ingestor() {
	queue.Init()

	chnMessages := make(chan *AWSremoveSQS.Message, MaxNumberOfMessages)
	go queue.PollMessages(chnMessages, MaxNumberOfMessages)

	for message := range chnMessages {
		handleMessage(message)

		if err := queue.DeleteMessage(*message.ReceiptHandle); err != nil {
			log.Errorf("failed to delete message %s", err)
		}
	}
}
