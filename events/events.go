package events

import (
	"fmt"

	"github.com/MNThomson/securitybot/pkg/queue"
	log "github.com/sirupsen/logrus"
)

func InitEvents() {
	queue.Init()

	message := "HELLO"

	err := queue.SendMsg(message)
	if err != nil {
		log.Errorf("Got an error sending the message: %s", err)
		return
	}

	fmt.Println("Sent message to queue ")
}
