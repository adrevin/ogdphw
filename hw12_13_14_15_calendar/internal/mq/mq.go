package mq

type MQ interface {
	SendEventNotification(notification *Notification) error
	Close()
}
