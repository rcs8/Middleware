package main

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

type ServerRabbitMQ struct {
	conn         *amqp.Connection
	channel      *amqp.Channel
	requestQueue *amqp.Queue
	sqrt         *SqrtRabbitMQ
}

func NewServerRabbitMQ(address string) (*ServerRabbitMQ, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()

	if err != nil {
		return nil, err
	}

	requestQueue, err := ch.QueueDeclare("request", false, false, false, false, nil)

	if err != nil {
		return nil, err
	}

	sqrt := new(SqrtRabbitMQ)

	return &ServerRabbitMQ{
		conn:         conn,
		channel:      ch,
		requestQueue: &requestQueue,
		sqrt:         sqrt,
	}, err
}

func (s *ServerRabbitMQ) ListenRabbitMQ(exit NotifChan, exited NotifChan) {
	msgsFromClient, err := s.channel.Consume(s.requestQueue.Name, "", true, false, false, false, nil)

	if err != nil {
		_, stop := <-exit
		if stop {
			s.conn.Close()
			s.channel.Close()
			exited <- true
			return
		}
		return
	}

	for d := range msgsFromClient {
		msgRequest := Request{}

		err := json.Unmarshal(d.Body, &msgRequest)

		if err != nil {
			_, stop := <-exit
			if stop {
				s.conn.Close()
				s.channel.Close()
				exited <- true
				return
			}
			continue
		}

		reply := s.sqrt.Sqrt(&msgRequest)

		replyMsgBytes, err := json.Marshal(reply)

		if err != nil {
			_, stop := <-exit
			if stop {
				s.conn.Close()
				s.channel.Close()
				exited <- true
				return
			}
			continue
		}

		err = s.channel.Publish("", d.ReplyTo, false, false, amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: d.CorrelationId,
			Body:          []byte(replyMsgBytes),
		})

		if err != nil {
			_, stop := <-exit
			if stop {
				s.conn.Close()
				s.channel.Close()
				exited <- true
				return
			}
			continue
		}
	}
}

func (s *ServerRabbitMQ) Close() {
	(*s.conn).Close()
	(*s.channel).Close()
}
