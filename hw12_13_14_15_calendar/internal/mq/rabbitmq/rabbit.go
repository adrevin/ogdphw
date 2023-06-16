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

func New(config configuration.MessageQueueConfiguration, logger logger.Logger) mq.MQ {
	conn, err := amqp091.Dial(config.BrokerURI)
	if err != nil {
		logger.Errorf("Failed to open connection: %+v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		logger.Errorf("Failed to get connection channel channel: %+v", err)
	}

	_, err = ch.QueueDeclare(config.QueueName, true, false, false, true, nil)
	if err != nil {
		logger.Fatalf("failed to connect to RabbitMQ: %+v", err)
	}

	return &rabbitMQ{config: config, logger: logger, conn: conn, ch: ch}
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
	/*	ch, err := r.conn.Channel()
		defer func(ch *amqp091.Channel) {
			err := ch.Close()
			if err != nil {
				r.logger.Errorf("Failed to close a channel: %+v", err)
			}
		}(ch)

		if err != nil {
			r.logger.Errorf("Failed to open a channel: %+v", err)
		}
	*/
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
	if err != nil {
		return err
	}

	return nil
}

/*
func (r *rabbitMQ) Consume(ctx context.Context, callback func(*mq.Notification) bool) error {
	ch, err := r.conn.Channel()

	delivery, err := ch.Consume(
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

	go func() {
		defer func(ch *amqp091.Channel) {
			err := ch.Close()
			if err != nil {
				r.logger.Errorf("Failed to close a channel: %+v", err)
			}
		}(ch)
		for {
			select {
			case <-ctx.Done():
				return
			case d := <-delivery:
				var notification *mq.Notification
				err = json.Unmarshal(d.Body, &notification)
				if err != nil {
					r.logger.Errorf("unmarshal error: v+%", err)
					continue
				}
				if callback(notification) {
					err := ch.Ack(d.DeliveryTag, false)
					if err != nil {
						r.logger.Errorf("can not ASK delivery: v+%", err)
					}
				} else {
					err := ch.Nack(d.DeliveryTag, false, true)
					if err != nil {
						r.logger.Errorf("can not NACK delivery: v+%", err)
					}
				}
			}
		}
	}()

	return nil
}
*/
