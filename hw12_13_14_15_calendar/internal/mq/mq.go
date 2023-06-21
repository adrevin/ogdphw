package mq

import "context"

type MQ interface {
	SendEventNotification(notification *Notification) error
	ConsumeNotifications(ctx context.Context, callback func(*Notification) bool) error
	Close()
}
