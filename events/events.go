package events

import (
	log "github.com/sirupsen/logrus"

	"github.com/MNThomson/securitybot/pkg/queue"
)

func InitEvents() {
	queue.Init()

	message := "HELLO"

	err := queue.SendMsg(message)
	if err != nil {
		log.Errorf("Got an error sending the message: %s", err)
		return
	}

	log.Debug("Sent message to queue")
}
