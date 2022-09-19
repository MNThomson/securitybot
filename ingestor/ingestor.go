package ingestor

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/MNThomson/securitybot/pkg/queue"
	"github.com/MNThomson/securitybot/pkg/types"
)

const (
	MaxNumberOfMessages int64 = 10
)

func handleMessage(alert types.AlertType) {
	log.Debug("RECEIVING MESSAGE >>> ")
	log.Debug(alert.Info)
}

func Ingestor(db *gorm.DB) {
	queue.Init()

	chnAlerts := make(chan types.AlertType, MaxNumberOfMessages)
	go queue.PollMessages(chnAlerts, MaxNumberOfMessages)

	for alert := range chnAlerts {
		handleMessage(alert)

		if err := queue.DeleteMessage(alert.ReceiptHandle); err != nil {
			log.Errorf("failed to delete message %s", err)
		}
	}
}
