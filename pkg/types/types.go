package types

type AlertType struct {
	ID            string
	OccuredAt     string
	ReceiptHandle string

	AlertNumber int64
	Actor       string
	Action      string
	Info        string
}
