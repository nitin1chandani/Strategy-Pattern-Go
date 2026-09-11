package lib

type NotificationType string

const (
	MAIL              NotificationType = "MAIL"
	SMS               NotificationType = "SMS"
	PUSH_NOTIFICATION NotificationType = "PUSH_NOTIFICATION"
)
