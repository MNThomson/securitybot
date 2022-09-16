package sqs

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var (
	sqsSvc          *sqs.SQS
	queueURL        *string
	queue           string = "SecurityAlertQueue"
	WaitTimeSeconds int64  = 15
	DelaySeconds    int64  = 0
)

func getQueueURL(sess *session.Session, queue *string) (*sqs.GetQueueUrlOutput, error) {
	svc := sqs.New(sess)

	urlResult, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: queue,
	})
	if err != nil {
		return nil, err
	}

	return urlResult, nil
}

func PollMessages(chn chan<- *sqs.Message, MaxNumberOfMessages int64) {

	for {
		output, err := sqsSvc.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl:            queueURL,
			MaxNumberOfMessages: &MaxNumberOfMessages,
			WaitTimeSeconds:     &WaitTimeSeconds,
		})
		if err != nil {
			fmt.Printf("failed to fetch sqs message %v", err)
		}

		for _, message := range output.Messages {
			chn <- message
		}
	}
}

func SendMsg(message string) error {
	_, err := sqsSvc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: &DelaySeconds,
		// MessageAttributes: map[string]*sqs.MessageAttributeValue{
		// 	"Title": {
		// 		DataType:    aws.String("String"),
		// 		StringValue: aws.String("The Whistler"),
		// 	},
		// 	"Author": {
		// 		DataType:    aws.String("String"),
		// 		StringValue: aws.String("John Grisham"),
		// 	},
		// 	"WeeksOn": {
		// 		DataType:    aws.String("String"),
		// 		StringValue: aws.String("6"),
		// 	},
		// },
		MessageBody: &message,
		QueueUrl:    queueURL,
	})
	if err != nil {
		return err
	}

	return nil
}

func DeleteMessage(ReceiptHandle string) {
	sqsSvc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      queueURL,
		ReceiptHandle: &ReceiptHandle,
	})
}

func InitSQS() {
	// Create a session that gets credentials
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Get URL of queue
	urlResult, err := getQueueURL(sess, &queue)
	if err != nil {
		fmt.Println("Got an error getting the queue URL:", err)
		return
	}

	queueURL = urlResult.QueueUrl
	sqsSvc = sqs.New(sess)
}
