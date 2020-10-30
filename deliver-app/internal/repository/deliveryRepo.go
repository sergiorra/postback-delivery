package repository

type DeliveryRepo interface {
	CheckConnection() error
	PopMessage() (string, error)
}