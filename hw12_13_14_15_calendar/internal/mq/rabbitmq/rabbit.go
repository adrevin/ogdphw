package rabbitmq

import (
	"context"
	"encoding/json"

	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/configuration"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/logger"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/mq"
	"github.com/rabbitmq/amqp091-go"
)

type rabbitMQ struct {
	config configuration.MessageQueueConfiguration
	logger logger.Logger
	conn   *amqp091.Connection
	ch     *amqp091.Channel
}

func New(config configuration.MessageQueueConfiguration, logger logger.Logger) (mq.MQ, error) {
	conn, err := amqp091.Dial(config.BrokerURI)
	if err != nil {
		logger.Errorf("failed to open RMQ connection: %+v", err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		logger.Errorf("failed to get RMQ connection channel channel: %+v", err)
		return nil, err
	}

	_, err = ch.QueueDeclare(config.QueueName, true, false, false, true, nil)
	if err != nil {
		logger.Errorf("failed to declare RMQ queue: %+v", err)
		return nil, err
	}

	return &rabbitMQ{config: config, logger: logger, conn: conn, ch: ch}, nil
}

func (r *rabbitMQ) Close() {
	err := r.ch.Close()
	if err != nil {
		r.logger.Errorf("channel close error: %+v", err)
	} else {
		r.logger.Debug("channel closed")
	}

	err = r.conn.Close()
	if err != nil {
		r.logger.Errorf("connection close error: %+v", err)
	} else {
		r.logger.Debug("connection closed")
	}
}

func (r *rabbitMQ) SendEventNotification(notification *mq.Notification) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.config.PublishTimeout)
	defer cancel()

	body, err := json.Marshal(notification)
	if err != nil {
		return err
	}

	err = r.ch.PublishWithContext(ctx, "", r.config.QueueName, false, false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	return err
}

func (r *rabbitMQ) ConsumeNotifications(ctx context.Context, callback func(*mq.Notification) bool) error {
	delivery, err := r.ch.Consume(
		r.config.QueueName, // queue
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	r.logger.Info("notifications consuming started")
	for {
		select {
		case <-ctx.Done():
			r.logger.Info("notifications consuming stopped")
			return nil
		case d := <-delivery:
			var notification *mq.Notification
			err = json.Unmarshal(d.Body, &notification)
			if err != nil {
				r.logger.Errorf("unmarshal error: v+%", err)
				continue
			}
			if callback(notification) {
				err := r.ch.Ack(d.DeliveryTag, false)
				if err != nil {
					r.logger.Errorf("can not ASK delivery: v+%", err)
				}
			} else {
				err := r.ch.Nack(d.DeliveryTag, false, true)
				if err != nil {
					r.logger.Errorf("can not NACK delivery: v+%", err)
				}
			}
		}
	}
}
